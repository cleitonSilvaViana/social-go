package repository

import (
	"database/sql"

	"github.com/cleitonSilvaViana/social-go/internal/database"
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

func (c *contactRepository) CreateNewUserContact(email string) (int, *fail.ResponseError) {
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
