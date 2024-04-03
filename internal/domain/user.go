package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cleitonSilvaViana/social-go/internal/repository"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type Data struct {
	FirstName string    `json="firstName" validate="required,lte=3"`
	BirthDate time.Time `json="birthDate" validate="required"`
}

type Address struct {
	Country string `json="country'`
	State   string `json="state"`
	City    string `json="city"`
}

type Credentials struct {
	Nick     string `json="nick" validate="required,lte=5"`
	Email    string `json="email" validate="required,email"`
	Password string `json="password" validate="required,lte=5"`
}

type Contact struct {
	Email string `json="email" validate="email"`
	Phone string `json="phone" validate=gt=13`
}

func CreateUser(bodyRequest []byte) error {
	type createUser struct {
		Data        `validate="required"`
		Address     `validate="required"`
		Credentials `validate="required"`
		Contact     `validade="required"`
	}

	var user createUser
	err := json.Unmarshal(bodyRequest, &user)
	if err != nil {
		return fail.ValidateFields(err)
	}

	repo, err := repository.NewUserRepository()
	if err != nil {
		return nil
	}

	err = repo.CreateUser()
	if err != nil {
		return err
	}

	return nil
}

func SearchUsers() ([]byte, error) {

	// adicionar lógica para ler query params

	repo, err := repository.NewUserRepository()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	users, err := repo.SearchUsers()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	JSONData, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return JSONData, nil
}
