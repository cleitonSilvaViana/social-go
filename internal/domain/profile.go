package domain

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/cleitonSilvaViana/social-go/internal/entitie"
	"github.com/cleitonSilvaViana/social-go/internal/repository"
	"github.com/cleitonSilvaViana/social-go/internal/security"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Credentials struct {
	Nick     string `json:"nick" validate:"required,lte=5"`
	Password string `json:"password" validate:"required,lte=5"`
}

func (c *Credentials) Trim() {
	c.Nick = strings.Trim(c.Nick, " ")
	c.Password = strings.Trim(c.Password, " ")
}

func (c *Credentials) Preper() {
	c.Nick = strings.ToLower(c.Nick)
}

func (c *Credentials) ValidateNick() *fail.ResponseError {
	if len(c.Nick) < 5 {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "o campo nick deve possuir ao menos 5 caracteres",
		}
	}

	// verificar se há caracteres especiais no nick
	// se houver, iremos retornar um erro

	return nil
}

func (c *Credentials) ValidatePassword() *fail.ResponseError {
	if len(c.Password) < 6 {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "a senha não pode possuir menos de 6 caracteres",
		}
	}
	return nil
}

func (c *Credentials) CheckIfNickExists() (bool, *fail.ResponseError) {

	repo, err := repository.NewProfileRepository()
	if err != nil {
		return false, err
	}

	param := map[string]string{
		"nick": c.Nick,
	}

	profile, err := repo.GetProfile(param)
	if err != nil {
		return false, err
	}

	return profile.Nick != "", nil
}

func (c *Credentials) Validate() *fail.ResponseError {
	var errs []error

	c.Trim()
	c.Preper()

	err := c.ValidateNick()
	if err != nil {
		errs = append(errs, err)
	}

	err = c.ValidatePassword()
	if err != nil {
		errs = append(errs, err)
	}

	_, err = c.CheckIfNickExists()
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		var msg string

		for _, e := range errs {
			msg += e.Error() + "\n"
		}

		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    msg,
		}
	}
	return nil
}

func Login(bodyRequest []byte) (string, *fail.ResponseError) {
	var c Credentials

	e := json.Unmarshal(bodyRequest, &c)
	if e != nil {
		return "", &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    fail.ValidateFields(e).Error(),
		}
	}

	err := c.Validate()
	if err != nil {
		return "", err
	}

	repo, err := repository.NewProfileRepository()
	if err != nil {
		return "", err
	}

	param := map[string]string{
		"nick": c.Nick,
	}

	profile, err := repo.GetProfile(param)
	if err != nil {
		return "", err
	}

	ok := security.ComparePasswords(profile.Password, []byte(c.Password))
	if !ok {
		// adicionar alguma mensagem de err
		return "", &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "usuário ou senha inválidos",
		}
	}

	currentDate := time.Now()
	expiration := currentDate.Add(time.Hour * 10)

	claims := security.UserClaims{
		Nick: c.Nick,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentDate.Unix(),
			ExpiresAt: expiration.Unix(),
		},
	}

	token, e := security.NewAccessToken(claims)
	if e != nil {

		// gerar log aqui
		slog.LogAttrs(context.Background(),
			slog.LevelError,
			"erro ao gerar token JWT:"+e.Error(),
		)

		return "", &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "internal server error",
		}
	}

	return token, nil
}

func Logout() {}

func CreateProfileUser(bodyRequest []byte) (string, *fail.ResponseError) {

	type newProfile struct {
		User
		Credentials
		Contact
	}

	var u newProfile

	e := json.Unmarshal(bodyRequest, &u)
	if e != nil {
		return "", &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    fail.ValidateFields(e).Error(),
		}
	}

	// tratando os dados
	ok := u.validateBirthDate()
	if !ok {
		return "", &fail.ResponseError{
			StatusCode: http.StatusUnauthorized,
			Message:    "apenas para maiores de 16 anos",
		}
	}

	// validar formato do nome, nick, email,
	


	// validando se dado já existe no sistema
	ok, err := u.CheckIfEmailExists()
	if err != nil {
		return "", err
	}

	if ok {
		return "", &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "email já está em uso",
		}
	}

	ok, err = u.CheckIfNickExists()
	if err != nil {
		return "", err
	}

	if ok {
		return "", &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "nick já está em uso",
		}
	}

	pass, err := security.CreateHashAndSalt(u.Password)
	if err != nil {
		return "", err
	}

	repoContact, err := repository.NewContactRepository()
	if err != nil {
		return "", err
	}

	contactID, err := repoContact.RegisterNewProfileEmail(u.Email)
	if err != nil {
		return "", err
	}

	repoProfile, err := repository.NewProfileRepository()
	if err != nil {
		return "", err
	}

	var p = entitie.Profile{
		UID:       uuid.NewString(),
		Nick:      u.Nick,
		Password:  pass,
		ContactID: contactID,
	}

	err = repoProfile.CreateProfileUser(p)
	if err != nil {
		return "", err
	}

	currentDate := time.Now()
	expiration := currentDate.Add(time.Hour * 10)

	claims := security.UserClaims{
		Nick: p.Nick,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentDate.Unix(),
			ExpiresAt: expiration.Unix(),
		},
	}

	token, e := security.NewAccessToken(claims)
	if e != nil {
		return "", &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "internal server error",
		}
	}

	return token, nil
}

func CreateProfileCompany(bodyRequest []byte) ([]byte, *fail.ResponseError) {
	return nil, nil
}


// refatorar
func SearchProfile(param map[string]string) ([]byte, *fail.ResponseError) {
	// os parametros enviados via query params podem ser de n tipos
	// buscar por um perfil cujo primeiro nome seja semelhante à por exemplo "maria" -> possíveis respostas -> mariana, marianna, etc.
	// buscar por perfis de pessoas cuja idade seja superior a 32 anos (query = "/users?age-tha=32") 




	repo, err := repository.NewProfileRepository()
	if err != nil {
		return nil, err
	}

	profile, err := repo.GetProfile(param)
	if err != nil {
		return nil, err
	}

	JSONData, e := json.Marshal(profile)
	if e != nil {
		return nil, &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    e.Error(),
		}
	}

	return JSONData, nil
}

func SearchProfiles(queryParams map[string]string) ([]byte, *fail.ResponseError) {
	var profiles *[]entitie.Profile

	repo, err := repository.NewProfileRepository()
	if err != nil {
		return nil, err
	}

	profiles, err = repo.GetProfiles(queryParams)
	if err != nil {
		return nil, err
	}

	JSONData, e := json.Marshal(*profiles)
	if e != nil {
		// criar log aqui
		return nil, fail.INTERNAL_SERVER_ERROR
	}
	return JSONData, nil
}
