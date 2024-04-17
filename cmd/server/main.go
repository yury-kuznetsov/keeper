package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	dataHandler "gophkeeper/internal/server/handler/data"
	userHandler "gophkeeper/internal/server/handler/user"
	"gophkeeper/internal/server/middleware"
	dataRepository "gophkeeper/internal/server/repository/data"
	userRepository "gophkeeper/internal/server/repository/user"
	dataService "gophkeeper/internal/server/service/data"
	userService "gophkeeper/internal/server/service/user"
	"gophkeeper/internal/server/service/webtoken"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("No .env file found")
	}

	// создаем сервер
	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: getRouter(),
	}

	// готовим канал для прослушивания системных сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// запускаем сервера в отдельной горутине
	go func() {
		log.Printf("Starting server on %s...\n", server.Addr)
		var err error
		httpsEnabled, _ := strconv.ParseBool(os.Getenv("HTTPS"))
		if httpsEnabled {
			err = server.ListenAndServeTLS(getCertAndKey())
		} else {
			err = server.ListenAndServe()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("HTTP server: %v", err)
		}
	}()

	// ожидаем сигнала остановки из канала `stop`
	<-stop

	// даем серверу 5 секунд на завершение обработки текущих запросов
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// завершаем "мягко" работу сервера
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("HTTP server Shutdown: %v", err)
	}
}

func getRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.GzipMiddleware)
	r.Use(chiMiddleware.Logger)

	db, err := sql.Open("pgx", os.Getenv("DB_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}

	// сервисы аутентификации
	userRepo := userRepository.NewRepository(db)
	userSvc := userService.NewService(userRepo)

	// сервис данных пользователя
	dataRepo := dataRepository.NewRepository(db)
	dataSvc := dataService.NewService(dataRepo)

	// сервис создания JWT
	jwtSvc := webtoken.NewJWTService()

	r.Post("/api/user/register", userHandler.RegisterHandler(userSvc, jwtSvc))
	r.Post("/api/user/login", userHandler.LoginHandler(userSvc, jwtSvc))

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(jwtSvc))
		r.Get("/api/data/pull", dataHandler.PullHandler(dataSvc))
		r.Post("/api/data/push", dataHandler.PushHandler(dataSvc))
		r.Get("/api/data/status", dataHandler.StatusHandler(dataSvc))
	})

	return r
}

func getCertAndKey() (certFile, keyFile string) {
	// создаём шаблон сертификата
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"GophKeeper"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	// создаём новый приватный RSA-ключ длиной 4096 бит
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatal(err)
	}

	// создаём сертификат x.509
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// кодируем сертификат в формате PEM
	var certPEM bytes.Buffer
	err = pem.Encode(&certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("cert.pem", certPEM.Bytes(), 0600)
	if err != nil {
		log.Fatalf("Failed to save certificate: %v", err)
	}

	// кодируем ключ в формате PEM
	var privateKeyPEM bytes.Buffer
	err = pem.Encode(&privateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("key.pem", privateKeyPEM.Bytes(), 0600)
	if err != nil {
		log.Fatalf("Failed to save private key: %v", err)
	}

	return "cert.pem", "key.pem"
}
