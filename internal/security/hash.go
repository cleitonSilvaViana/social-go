package security

import (
	"net/http"

	"github.com/cleitonSilvaViana/social-go/pkg/fail"
	"golang.org/x/crypto/bcrypt"
)

// CreateHashAndSalt is responsibility for create a hash of password of user
func CreateHashAndSalt(password string) ([]byte, *fail.ResponseError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
			// criar um log neste ponto
			// Um erro ocorrido neste trecho de código é sensível e não deve ser exposto ao cliente
			// iremos retornar um erro generico ao cliente
			// para o dev, iremos criar um log relatando o erro específico
		return nil, &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message: "internal server error",
		}
	}
	return hash, nil
}

// ComparePassword is used for compare the passwords
func ComparePasswords(hashedPassord, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassord, password)
	return err == nil
}
