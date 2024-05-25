package middleware

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

// GzipMiddleware is a middleware function that handles gzip compression for HTTP requests.
// It checks if the client sent compressed data and decompresses it if necessary.
// It also checks if the client can accept compressed data and compresses the response if necessary.
// The function returns an http.Handler that handles the middleware logic.
func GzipMiddleware(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// проверяем, что клиент отправил сжатые данные
			contentEncoding := r.Header.Get("Content-Encoding")
			if strings.Contains(contentEncoding, "gzip") {
				gr, err := gzip.NewReader(r.Body)
				if err != nil {
					fmt.Println("ERROR: " + err.Error())
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				defer func(gr *gzip.Reader) {
					_ = gr.Close()
				}(gr)
				r.Body = gr
			}

			// проверяем, что клиент умеет принимать сжатые данные
			acceptEncoding := r.Header.Get("Accept-Encoding")
			if strings.Contains(acceptEncoding, "gzip") {
				gw := NewGzipWriter(w)
				defer gw.Close()
				next.ServeHTTP(gw, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}(next)
}

// GzipWriter is a type that represents a gzip writer object
type GzipWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

// NewGzipWriter creates a new GzipWriter object that wraps around an http.ResponseWriter.
// It initializes a gzip.Writer and sets "Content-Encoding" header to "gzip" in the response.
// The function returns a pointer to the GzipWriter object.
func NewGzipWriter(w http.ResponseWriter) *GzipWriter {
	gw := gzip.NewWriter(w)
	w.Header().Set("Content-Encoding", "gzip")
	return &GzipWriter{gw, w}
}

// Write is a method of the GzipWriter type that writes the given byte slice to the underlying gzip writer.
// It returns the number of bytes written and any error encountered during the write operation.
// The method delegates the write operation to the underlying gzip.Writer object and returns its result.
// The method signature is:
//
//	func (w *GzipWriter) Write(b []byte) (int, error)
//
// Example usage:
//
//	gw := NewGzipWriter(w)
//	defer gw.Close()
//	_, err := gw.Write([]byte("Hello, world!"))
//	if err != nil {
//	  fmt.Println("Error writing data:", err)
//	}
//
// Note: The Close method of the GzipWriter should be called after the write operation is complete to flush any buffered data.
//
//	Failure to do so may result in incomplete or corrupted output.
func (w *GzipWriter) Write(b []byte) (int, error) {
	return w.gw.Write(b)
}

// Close is a method of the GzipWriter type that closes the underlying gzip writer.
// It flushes any buffered data and releases any resources associated with the writer.
//
// Example usage:
//
//	gw := NewGzipWriter(w)
//	defer gw.Close()
//	_, err := gw.Write([]byte("Hello, world!"))
//	if err != nil {
//	  fmt.Println("Error writing data:", err)
//	}
//
// Note: The Close method of the GzipWriter should be called after the write operation is complete to flush any buffered data.
//
//	Failure to do so may result in incomplete or corrupted output.
func (w *GzipWriter) Close() {
	_ = w.gw.Close()
}
