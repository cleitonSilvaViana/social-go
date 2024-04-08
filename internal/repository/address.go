package repository

import (
	"database/sql"

	"github.com/cleitonSilvaViana/social-go/internal/database"
	"github.com/cleitonSilvaViana/social-go/internal/entities"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type addressRepository struct {
	db *sql.DB
}

func NewAddressRepository() (*addressRepository, *fail.ResponseError) {
	db, err := database.ConnectToMysql()
	if err != nil {
		return nil, fail.NewMySqlError(err)
	}
	return &addressRepository{db}, nil
}

func (a *addressRepository) GetCountryCode(country string) (string, *fail.ResponseError) {
	defer a.db.Close()

	row := a.db.QueryRow("SELECT cca3, name FROM country WHERE name = ?", country)

	var c entities.Country

	err := row.Scan(
		&c.CCA3,
		&c.Name,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", nil
		} else {
			return "", fail.NewMySqlError(err)
		}
	}

	return c.CCA3, nil
}

func (a *addressRepository) AddNewCountry(country entities.Country) (string, *fail.ResponseError) {
	defer a.db.Close()

	stmt, err := a.db.Prepare("INSERT INTO country (cca3, name) VALUES (?, ?)")
	if err != nil {
		return "", fail.NewMySqlError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(country.CCA3, country.Name)
	if err != nil {
		return "", fail.NewMySqlError(err)
	}

	return country.CCA3, nil
}

func (a *addressRepository) GetStateID(state, countryCCA3 string) (int, *fail.ResponseError) {
	defer a.db.Close()

	row := a.db.QueryRow("SELECT id FROM state WHERE countryID = ? AND name = ?", countryCCA3, state)

	var id int

	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	return id, nil
}

func (a *addressRepository) AddNewState(state, countryCCA3 string) (int, *fail.ResponseError) {
	defer a.db.Close()

	stmt, err := a.db.Prepare("INSERT INTO state (name, countryID) VALUES (?, ?)",)
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(state, countryCCA3)
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	stateID, err := result.LastInsertId()
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	return int(stateID), nil
}

func (a *addressRepository) GetCityID(city string, stateID int) (int, *fail.ResponseError) {
	defer a.db.Close()

	row := a.db.QueryRow("SELECT id FROM city WHERE name = ?", city)

	var cityID int
	
	err := row.Scan(&cityID)
	if err == sql.ErrNoRows {
		return 0, nil
	} 

	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	return cityID, nil
}

func (a *addressRepository) AddNewCity(city string, stateID int) (int, *fail.ResponseError) {
	defer a.db.Close()

	stmt, err := a.db.Prepare("INSERT INTO city (name, stateID) VALUES (?, ?)",)
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(city, stateID)
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	cityID, err := result.LastInsertId()
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	return int(cityID), nil
}