package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cleitonSilvaViana/social-go/internal/database"
	"github.com/cleitonSilvaViana/social-go/internal/entitie"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type profielRepository struct {
	db *sql.DB
}

func NewProfileRepository() (*profielRepository, *fail.ResponseError) {
	db, err := database.ConnectToMysql()
	if err != nil {
		return nil, fail.INTERNAL_SERVER_ERROR
	}
	return &profielRepository{db}, nil
}

func (p *profielRepository) CreateProfileUser(profile entitie.Profile) *fail.ResponseError {

	defer p.db.Close()

	stmt, err := p.db.Prepare("INSERT INTO user (uid, nick, password) VALUES (?, ?, ?)")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(profile.UID, profile.Nick, profile.Password)
	if err != nil {
		return fail.NewMySqlError(err)
	}

	return nil
}

func (p *profielRepository) CreateProfileCompany(profile entitie.Profile) *fail.ResponseError {
	return nil
}

func (p *profielRepository) UpdateNickProfile(UID []byte, newNick string) *fail.ResponseError {
	defer p.db.Close()

	stmt, err := p.db.Prepare("UPDATE profile SET nick = ? WHERE uid = ?")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	_, err = stmt.Exec(newNick, UID)
	if err != nil {
		return fail.NewMySqlError(err)
	}

	return nil
}


func (p *profielRepository) UpdatePasswordProfile(UID, newPassword []byte) *fail.ResponseError {
	defer p.db.Close()

	stmt, err := p.db.Prepare("UPDATE profile SET password = ? WHERE uid = ?")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	_, err = stmt.Exec(newPassword, UID)
	if err != nil {
		return fail.NewMySqlError(err)
	}

	return nil
}

func (p *profielRepository) GetProfile(param map[string]string) (*entitie.Profile, *fail.ResponseError) {
	defer p.db.Close()

	var (
		key string
		value string
	)

	for k, v := range param {
		key = k + " = ?"
		value = v
	}

	row := p.db.QueryRow("SELECT uid, nick, password, createdAt, contactID, addressID FROM profile WHERE ? = ?", key, value)

	var profile entitie.Profile

	err := row.Scan(&profile)

	if err != nil && err != sql.ErrNoRows {
		return nil, fail.NewMySqlError(err)
	}

	return &profile, nil
}





// refatorar
// maior que, menor que, igual à, diferente de, semelhante à...
func (p *profielRepository) GetProfiles(filters map[string]string) (*[]entitie.Profile, *fail.ResponseError) {
	defer p.db.Close()

	var keys string
	var values string

	for key, value := range filters {
		keys += " " + key + " = ?"
		values += " " + value
	}

	query := fmt.Sprint(`SELECT uid, nick, password, createdAt, contactID, addressID 
												FROM profile 
												WHERE %s`, keys)

	rows, err := p.db.Query(query, values)

	if err != nil {
		// gerar log de erro aqui
		return nil, fail.INTERNAL_SERVER_ERROR
	}

	defer rows.Close()

	var profiles []entitie.Profile

	for rows.Next() {
		var profile entitie.Profile

		err = rows.Scan(
			&profile.UID,
			&profile.Nick,
			&profile.Password,
			&profile.CreatedAt,
			&profile.ContactID,
			&profile.AddressID,
		)
		if err != nil {
			// gerar log...

			return nil, &fail.ResponseError{
				StatusCode: http.StatusInternalServerError,
				Message: "internal server error",
			}
		}

		profiles = append(profiles, profile)

	}

	return &profiles, nil
}