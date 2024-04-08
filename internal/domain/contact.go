package domain

import (
	"net/mail"

	"github.com/cleitonSilvaViana/social-go/internal/repository"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

func CreateNewUsercontact(email string) (int, *fail.ResponseError) {

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

	contactID, err := repo.CreateNewUserContact(email)
	if err != nil {
		return 0, err
	}

	return contactID, nil
}
