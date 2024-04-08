package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/cleitonSilvaViana/social-go/internal/entities"
	"github.com/cleitonSilvaViana/social-go/internal/repository"
	"github.com/cleitonSilvaViana/social-go/internal/security"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"

	"github.com/google/uuid"
)

type Data struct {
	UID       []byte    `json:"uid,omitempty"`
	FirstName string    `json="firstName" validate="required,lte=3"`
	BirthDate time.Time `json="birthDate" validate="required"`
}

func (d *Data) validateFirstName() *fail.ResponseError {
	d.FirstName = strings.Trim(d.FirstName, " ")
	if len(d.FirstName) < 3 {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "seu nome não pode conter menos de 3 caracteres",
		}
	}
	return nil
}

// validateBirthDate return true case the user is older than 16 years old
func (d *Data) validateBirthDate() bool {
	now := time.Now()
	minimunAge := now.Add(-time.Duration(now.Year()))
	return d.BirthDate.Before(minimunAge)
}

type Address struct {
	Country string `json="country'`
	State   string `json="state"`
	City    string `json="city"`
}

type Credentials struct {
	Nick     string `json:"nick" validate:"required,lte=5"`
	Password string `json:"password" validate:"required,lte=5"`
}

func (c *Credentials) validateNick() *fail.ResponseError {
	c.Nick = strings.Trim(c.Nick, " ")
	if len(c.Nick) < 5 {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "o campo nick deve possuir ao menos 5 caracteres",
		}
	}
	return nil
}

func (c *Credentials) validatePassword() *fail.ResponseError {
	c.Password = strings.Trim(c.Password, " ")
	if len(c.Password) < 6 {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "a senha não pode possuir menos de 6 caracteres",
		}
	}
	return nil
}

type Contact struct {
	Phone string `json:"phone"`
	Email string `json:"email" validate:"required,email"`
}

func (c *Contact) ValidateEmail() *fail.ResponseError {
	c.Email = strings.Trim(c.Email, " ")
	if c.Email == "" {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "o campo de email não pode estar vazio",
		}
	}

	_, err := mail.ParseAddress(c.Email)
	if err != nil {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "o email não está em um formato válido",
		}
	}
	return nil
}

func (c *Contact) ValidatePhone() error {
	return nil
}

type User struct {
	Data
	Address
	Credentials
	Contact
}

func (u *User) validade() *fail.ResponseError {

	ok := u.validateBirthDate()
	if !ok {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "apenas maiores de 16 anos podem se cadastrar",
		}
	}

	err := u.ValidateEmail()
	if err != nil {
		return err
	}

	err = u.validatePassword()
	if err != nil {
		return err
	}

	err = u.validateNick()
	if err != nil {
		return err
	}

	err = u.validateFirstName()
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(bodyRequest []byte) *fail.ResponseError {

	var u User
	e := json.Unmarshal(bodyRequest, &u)
	if e != nil {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    fail.ValidateFields(e).Error(),
		}
	}

	err := u.validade()
	if err != nil {
		return err
	}

	pass, err := security.CreateHashAndSalt(u.Password)
	if err != nil {
		return err
	}

	countryCode, err := GetCountryDatabase(u.Country)
	if err != nil {
		return err
	}

	stateID, err := GetState(u.State, countryCode)
	if err != nil {
		return err
	}

	cityID, err := GetCityID(u.City, stateID)
	if err != nil {
		return err
	}

	contactID, err := CreateNewUsercontact(u.Email)
	if err != nil {
		return err
	}

	var user = entities.User{
		UID:       uuid.NewString(),
		Password:  pass,
		CityID:    cityID,
		Nick:      u.Nick,
		FirstName: u.FirstName,
		BirthDate: u.BirthDate,
		ContactID: contactID,
	}

	repo, err := repository.NewUserRepository()
	if err != nil {

		slog.LogAttrs(
			context.Background(),
			slog.LevelError,
			"error to connect MySQL",
			slog.String("error", err.Error()),
		)

		return err
	}

	err = repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func SearchUsers() ([]byte, *fail.ResponseError) {

	// adicionar lógica para ler query params

	repo, err := repository.NewUserRepository()
	if err != nil {
		return nil, err
	}

	users, err := repo.SearchUsers()
	if err != nil {
		return nil, err
	}

	JSONData, e := json.Marshal(*users)
	if e != nil {
		// criar log aqui
		return nil, fail.INTERNAL_SERVER_ERROR
	}

	return JSONData, nil
}

func SearchUser(UID string) ([]byte, *fail.ResponseError) {

	repo, err := repository.NewUserRepository()
	if err != nil {
		return nil, err
	}

	user, err := repo.SearchUserByID([]byte(UID))
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	JSONData, e := json.Marshal(user)
	if e != nil {
		return nil, &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    e.Error(),
		}
	}

	return JSONData, nil
}
