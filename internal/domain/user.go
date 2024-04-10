package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/cleitonSilvaViana/social-go/internal/entitie"
	"github.com/cleitonSilvaViana/social-go/internal/repository"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"

	"github.com/google/uuid"
)

type User struct {
	UID       []byte    `json:"uid,omitempty"`
	FirstName string    `json:"firstName" validate:"required,lte=3"`
	LastName  string    `json:"lastName" validate:"required,lte=3"`
	BirthDate time.Time `json:"birthDate" validate:"required"`
	Gender    string    `json:"gender"`
}

func (u *User) validateFirstName() *fail.ResponseError {
	u.FirstName = strings.Trim(u.FirstName, " ")
	if len(u.FirstName) < 3 {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "seu nome não pode conter menos de 3 caracteres",
		}
	}
	return nil
}

func (u *User) ValidateLastName() *fail.ResponseError {
	u.LastName = strings.Trim(u.LastName, " ")
	if len(u.LastName) < 3 {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "seu sobrenome não pode conter menos de 3 caracteres",
		}
	}
	return nil
}

// validateBirthDate return true case the user is older than 16 years old
func (u *User) validateBirthDate() bool {
	now := time.Now()
	minimunAge := now.Add(-time.Duration(now.Year()))
	return u.BirthDate.Before(minimunAge)
}

func (u *User) validateGender() *fail.ResponseError {
	u.Gender = strings.Trim(u.Gender, " ")
	if u.Gender == "" {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "o campo de genero deve estar preenchido",
		}
	}
	return nil
}

func (u *User) validade() *fail.ResponseError {
	var errs []error

	ok := u.validateBirthDate()
	if !ok {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "apenas maiores de 16 anos podem cadastrar-se nesta plataforma",
		}
	}

	err := u.validateFirstName()
	if err != nil {
		errs = append(errs, err)
	}

	err = u.ValidateLastName()
	if err != nil {
		errs = append(errs, err)
	}


	if len(errs) > 0 {
		var message string

		for _, err := range errs {
			message = fmt.Sprint("%s \n", err.Error())
		}

		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    message,
		}
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


	var user = entitie.User{
		UID:       uuid.NewString(),
		FirstName: u.FirstName,
		LastName: u.LastName,
		BirthDate: u.BirthDate,
		Gender: u.Gender,
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

func CreateUser2() *fail.ResponseError {
return nil
}