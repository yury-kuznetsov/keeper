package data

import (
	"encoding/json"
	"flag"
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// CreateHandler takes a Service and an []string of arguments, and performs the desired action based on the first argument.
// It returns an error if any error occurs during the execution of the desired action.
func CreateHandler(svc Service, arguments []string) error {
	var err error

	switch arguments[0] {

	case "credentials":
		_, err = createCredentials(svc, arguments[1:])

	case "text":
		_, err = createText(svc, arguments[1:])

	case "binary":
		_, err = createBinary(svc, arguments[1:])

	case "card":
		_, err = createCard(svc, arguments[1:])

	default:
		flag.PrintDefaults()
		return nil
	}

	return err
}

func createCredentials(svc Service, arguments []string) (uuid.UUID, error) {
	fs := flag.NewFlagSet("credentials", flag.ExitOnError)
	login := fs.String("l", "", "login")
	password := fs.String("p", "", "password")

	err := fs.Parse(arguments)
	if err != nil {
		return uuid.Nil, err
	}

	data := model.DataCredentials{
		Login:    *login,
		Password: *password,
	}
	dataJSON, _ := json.Marshal(data)

	return svc.Create(model.CategoryCredentials, dataJSON)
}

func createText(svc Service, arguments []string) (uuid.UUID, error) {
	fs := flag.NewFlagSet("text", flag.ExitOnError)
	text := fs.String("t", "", "text")

	err := fs.Parse(arguments)
	if err != nil {
		return uuid.Nil, err
	}

	data := model.DataText{Text: *text}
	dataJSON, _ := json.Marshal(data)

	return svc.Create(model.CategoryText, dataJSON)
}

func createBinary(svc Service, arguments []string) (uuid.UUID, error) {
	fs := flag.NewFlagSet("binary", flag.ExitOnError)
	binary := fs.String("b", "", "binary data")

	err := fs.Parse(arguments)
	if err != nil {
		return uuid.Nil, err
	}

	data := model.DataBinary{Binary: []byte(*binary)}
	dataJSON, _ := json.Marshal(data)

	return svc.Create(model.CategoryBinary, dataJSON)
}

func createCard(svc Service, arguments []string) (uuid.UUID, error) {
	fs := flag.NewFlagSet("card", flag.ExitOnError)
	number := fs.String("n", "", "number")
	owner := fs.String("o", "", "owner")
	cvv := fs.Uint("c", 0, "cvv")

	err := fs.Parse(arguments)
	if err != nil {
		return uuid.Nil, err
	}

	data := model.DataCard{
		Number: *number,
		Owner:  *owner,
		CVV:    *cvv,
	}
	dataJSON, _ := json.Marshal(data)

	return svc.Create(model.CategoryCard, dataJSON)
}
