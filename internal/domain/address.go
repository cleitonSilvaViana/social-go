package domain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cleitonSilvaViana/social-go/internal/entitie"
	"github.com/cleitonSilvaViana/social-go/internal/repository"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type Country struct {
	CountryCCA3 string `countryCca3,omitempty`
	Name        string `name,required`
}

func (c *Country) validateNameCountry() *fail.ResponseError {
	// realizar uma requisição http ao site que resgato infos sobre os países
	// checar se o país com nome equiavalente ao campo Name da struct existe

	return nil
}

type State struct {
	// padronização ISO
	Name string `json:"name,required"`
}

func (s *State) validate() *fail.ResponseError {
		s.Name = strings.Trim(s.Name, " ")

	if s.Name == "" {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message: "o nome do estado não pode estar vazio",
		}
	}

	return nil
}

type City struct {
	Name string `json:"city,required"`
}

func (c *City) validate() *fail.ResponseError {
	c.Name = strings.Trim(c.Name, " ")

	if c.Name == "" {
		return &fail.ResponseError{
			StatusCode: http.StatusBadRequest,
			Message: "o nome do estado não pode estar vazio",
		}
	}

	return nil
}

type Address struct {
	Country
	State
	City
}

func (a *Address) ValidateState() *fail.ResponseError {
	// verificar se o estado correspondente ao países celecionado existe

	return nil
}

func (a *Address) ValidateCity() *fail.ResponseError {
	// verificar se a cidade existe em determinado estado
	
	return nil
}


func (a *Address) Validate() *fail.ResponseError {
	// var errs []error
	
	a.ValidateState()
	a.ValidateCity()

	return nil
}

// SearchCountry irá realizar realizar uma busca no banco de dados pelo país inserido em sue parãmetro.
// Vale ressaltar que o país deve estar escrito em inglês
//
// Param:
//
//	country é o nome do país que será pesquisado no banco de dados
//
// Return:
//
//	a string retornada é o código do pais no formato padronizado pela ISO
//	error como o próprio nome diz, é um erro ocorrido durante a execução da função
func GetCountryDatabase(country string) (string, *fail.ResponseError) {

	repo, err := repository.NewAddressRepository()
	if err != nil {
		return "", err
	}

	code, err := repo.GetCountryCode(country)
	if err != nil {
		return "", err
	}

	if code == "" {
		code, err := AddNewContryInDatabase(country)
		if err != nil {
			return "", err
		}
		return code, nil
	}

	return code, nil
}

func GetCountryAPI(countryName string) (*Country, *fail.ResponseError) {

	const API_COUNTRIES = "https://restcountries.com/v3.1/name/"

	var client http.Client

	request, err := http.NewRequest("GET", API_COUNTRIES+countryName, nil)
	if err != nil {

		// criar um log aqui...

		return nil, &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

	response, err := client.Do(request)
	if err != nil {
			// criar um log aqui...

		return nil, &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
			// criar um log aqui...

		return nil, &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

	type JSON struct {
		Name struct {
			Common string `json:"common"`
		} `json:"name"`
		CCA3 string `json:"cca3"`
	}

	var datasJSON JSON

	bodyString := string(body)
	bodyString = strings.Trim(bodyString, "[")
	bodyString = strings.Trim(bodyString, "]")

	body = []byte(bodyString)

	err = json.Unmarshal(body, &datasJSON)
	if err != nil {
		// log de erro ao descerializar dados

		return nil, &fail.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

	return &Country{
		CountryCCA3: datasJSON.CCA3,
		Name:        datasJSON.Name.Common,
	}, nil
}

func AddNewContryInDatabase(country string) (string, *fail.ResponseError) {
	newCountry, err := GetCountryAPI(country)
	if err != nil {
		return "", err
	}

	repo, err := repository.NewAddressRepository()
	if err != nil {
		return "", err
	}

	code, err := repo.AddNewCountry(entitie.Country{
		CCA3: newCountry.CountryCCA3,
		Name: newCountry.Name,
	})
	if err != nil {
		return "", err
	}

	return code, nil
}

func GetState(state, countryCCA3 string) (int, *fail.ResponseError) {
	repo, err := repository.NewAddressRepository()
	if err != nil {
		return 0, err
	}

	stateID, err := repo.GetStateID(state, countryCCA3)
	if err != nil {
		return 0, err
	}

	if stateID == 0 {
		stateID, err = AddNewStateInDatabase(state, countryCCA3)
		if err != nil {
			return 0, err
		}
	}

	fmt.Println(stateID)

	return stateID, nil
}

func AddNewStateInDatabase(state, countryCCA3 string) (int, *fail.ResponseError) {
	repo, err := repository.NewAddressRepository()
	if err != nil {
		return 0, err
	}

	stateID, err := repo.AddNewState(state, countryCCA3)
	if err != nil {
		return 0, err
	}

	return stateID, nil
}

func GetCityID(city string, stateID int) (int, *fail.ResponseError) {
	repo, err := repository.NewAddressRepository()
	if err != nil {
		return 0, err
	}

	cityID, err := repo.GetCityID(city, stateID)
	if err != nil {
		return 0, err
	}

	if cityID == 0 {
		cityID, err = AddNewCityInDatabase(city, stateID)
		if err != nil {
			return 0, err
		}
	}

	return cityID, nil
}

func AddNewCityInDatabase(city string, stateID int) (int, *fail.ResponseError) {
	repo, err := repository.NewAddressRepository()
	if err != nil {
		return 0, err
	}

	cityID, err := repo.AddNewCity(city, stateID)
	if err != nil {
		return 0, err
	}

	return cityID, nil
}
