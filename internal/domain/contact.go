package domain

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/cleitonSilvaViana/social-go/internal/repository"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type Contact struct {
	Phone string `json:"phone"`
	Email string `json:"email" validate:"required,email"`
}

func (c *Contact) Trim() {
	c.Email = strings.Trim(c.Email, " ")
	c.Phone = strings.Trim(c.Phone, " ")
}

func (c *Contact) Preper() {
	c.Email = strings.ToLower(c.Email)
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

func (c *Contact) ValidatePhone() *fail.ResponseError {
	// verificar se há letras e/ou caracteres especiais no endereco telefônico
	// caso haja, retornaremos um erro...
	return nil
}

func (c *Contact) CheckIfEmailExists() (bool, *fail.ResponseError) {

	repo, err := repository.NewContactRepository()
	if err != nil {
		return false, err
	}
	
	var param = map[string]string{
		"email": c.Email,
	}

	contact, err := repo.GetContact(param)
	if err != nil {
		return false, err
	}

	return contact.Id > 0, nil
}

func (c *Contact) CheckIfPhoneExistis() (bool, *fail.ResponseError) {
	// verificar se existe o número fornecido já possuí registro no sistema
	return false, nil
}

/******************************************************************/
/******************************************************************/
/******************************************************************/

func (c *Contact) Validate() (*fail.ResponseError) {
	var errs []error
	
	c.Trim()
	c.Preper()

	c.ValidateEmail()
	c.CheckIfEmailExists()

	if c.Phone != "" {
		c.ValidatePhone()
		c.CheckIfPhoneExistis()
	}

	if len(errs) > 0 {
		var msg string
		
		for _, err := range errs {
			msg += err.Error() + "\n"
		}

		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message: msg,
		}
	}

	return nil
}

/******************************************************************/
/******************************************************************/
/******************************************************************/

func CreateNewProfilecontact(email string) (int, *fail.ResponseError) {

	_, e := mail.ParseAddress(email)
	if e != nil {

		// criar log aqui...

		return 0, fail.INTERNAL_SERVER_ERROR
	}

	// Verificando se o email está em um domínio válido
	// Verificando se o email exite e é utilizado
	// iremos enviar um e-mail com código randômico ao usuário

	repo, err := repository.NewContactRepository()
	if err != nil {
		return 0, err
	}

	contactID, err := repo.RegisterNewProfileEmail(email)
	if err != nil {
		return 0, err
	}

	return contactID, nil
}
