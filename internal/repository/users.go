package repository

import (
	"database/sql"
	"fmt"

	"github.com/cleitonSilvaViana/social-go/internal/database"
	"github.com/cleitonSilvaViana/social-go/internal/entitie"
	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type user struct {
	First_name string
	Last_name  string
	Nick       string
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() (*userRepository, *fail.ResponseError) {
	db, err := database.ConnectToMysql()
	if err != nil {
		return nil, fail.INTERNAL_SERVER_ERROR
	}
	return &userRepository{db}, nil
}

func (u *userRepository) CreateUser(user entitie.User) *fail.ResponseError {

	defer u.db.Close()

	stmt, err := u.db.Prepare("INSERT INTO USER (uid, first_name, last_name, gender, birth_date) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.UID, user.FirstName, user.LastName, user.Gender, user.BirthDate,
	)
	if err != nil {
		return fail.NewMySqlError(err)
	}

	return nil
}

func (u *userRepository) DeleteUser(UID string) *fail.ResponseError {
	defer u.db.Close()

	fmt.Println(UID)

	stmt, err := u.db.Prepare("DELETE FROM user WHERE uid = ?")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(UID)
	if err != nil {
		return fail.NewMySqlError(err)
	}

	return nil
}
