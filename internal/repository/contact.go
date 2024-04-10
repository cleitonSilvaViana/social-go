package repository

import (
	"database/sql"

	"github.com/cleitonSilvaViana/social-go/internal/database"
	"github.com/cleitonSilvaViana/social-go/internal/entitie"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository() (*contactRepository, *fail.ResponseError) {
	db, err := database.ConnectToMysql()
	if err != nil {
		return nil, fail.INTERNAL_SERVER_ERROR
	}
	return &contactRepository{db}, nil
}

func (c *contactRepository) RegisterNewProfileEmail(email string) (int, *fail.ResponseError) {
	defer c.db.Close()

	stmt, err := c.db.Prepare("INSERT INTO contact (email) VALUES (?)")
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	result, err := stmt.Exec(email)
	if err != nil {
		return 0, fail.NewMySqlError(err)
	}

	contactID, err := result.LastInsertId()
	if err != nil {
		return 0, fail.INTERNAL_SERVER_ERROR
	}

	return int(contactID), nil
}

func (c *contactRepository) RegisterNewProfilePhone(phone string) (int, *fail.ResponseError) {
	return 0, nil
}

// esboço...
// vou receber um par chave/valor que será utilizado para realizar a query
func (c *contactRepository) GetContact(param map[string]string) (*entitie.Contact, *fail.ResponseError) {
	defer c.db.Close()

	// basicamente, a chave do map passado como parâmetro deverá ser extritamente igual ao nome da coluna que será utilizada para filtrar os resultados
	// já seu respectivo valor, será o parâmetro da consulta
	row := c.db.QueryRow("SELECT id, email, phone FROM contact WHERE ? = ?", param /* chave */, param /* valor */ )
	
	var contact entitie.Contact

	err := row.Scan(&contact)
	if err != nil && err != sql.ErrNoRows{
		return nil, fail.NewMySqlError(err)
	}

	return &contact, nil
}