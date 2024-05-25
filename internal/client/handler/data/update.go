package data

import (
	"encoding/json"
	"flag"
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// UpdateHandler handles updating a data record based on the provided arguments.
// It takes a Service interface and a slice of arguments as input.
// The function switches on the first argument to determine which type of update to perform.
// It then calls the corresponding update function passing the Service interface and the remaining arguments.
// If the provided argument does not match any case, it prints the default flag options and returns nil.
// The function returns any error encountered during the update process.
func UpdateHandler(svc Service, arguments []string) error {
	var err error

	switch arguments[0] {

	case "credentials":
		err = updateCredentials(svc, arguments[1:])

	case "text":
		err = updateText(svc, arguments[1:])

	case "binary":
		err = updateBinary(svc, arguments[1:])

	case "card":
		err = updateCard(svc, arguments[1:])

	default:
		flag.PrintDefaults()
		return nil
	}

	return err
}

func updateCredentials(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("credentials", flag.ExitOnError)
	id := fs.String("i", "", "id")
	login := fs.String("l", "", "login")
	password := fs.String("p", "", "password")

	err := fs.Parse(arguments)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(*id)
	if err != nil {
		return err
	}

	data := model.DataCredentials{
		Login:    *login,
		Password: *password,
	}
	dataJSON, _ := json.Marshal(data)

	return svc.Update(uid, dataJSON)
}

func updateText(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("text", flag.ExitOnError)
	id := fs.String("i", "", "id")
	text := fs.String("t", "", "text")

	err := fs.Parse(arguments)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(*id)
	if err != nil {
		return err
	}

	data := model.DataText{Text: *text}
	dataJSON, _ := json.Marshal(data)

	return svc.Update(uid, dataJSON)
}

func updateBinary(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("binary", flag.ExitOnError)
	id := fs.String("i", "", "id")
	binary := fs.String("b", "", "binary data")

	err := fs.Parse(arguments)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(*id)
	if err != nil {
		return err
	}

	data := model.DataBinary{Binary: []byte(*binary)}
	dataJSON, _ := json.Marshal(data)

	return svc.Update(uid, dataJSON)
}

func updateCard(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("card", flag.ExitOnError)
	id := fs.String("i", "", "id")
	number := fs.String("n", "", "number")
	owner := fs.String("o", "", "owner")
	cvv := fs.Uint("c", 0, "cvv")

	err := fs.Parse(arguments)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(*id)
	if err != nil {
		return err
	}

	data := model.DataCard{
		Number: *number,
		Owner:  *owner,
		CVV:    *cvv,
	}
	dataJSON, _ := json.Marshal(data)

	return svc.Update(uid, dataJSON)
}
