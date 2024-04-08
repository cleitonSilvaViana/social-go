package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cleitonSilvaViana/social-go/internal/database"
	"github.com/cleitonSilvaViana/social-go/internal/entities"
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

func (u *userRepository) CreateUser(user entities.User) *fail.ResponseError {

	defer u.db.Close()

	stmt, err := u.db.Prepare("INSERT INTO USER (uid, nick, first_name, birth_date, password, cityID, contactID) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.UID, user.Nick, user.FirstName, user.BirthDate, user.Password, user.CityID, user.ContactID,
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

func (u *userRepository) AlterPassword(password []byte, uid []byte) *fail.ResponseError {
	defer u.db.Close()

	stmt, err := u.db.Prepare("UPDATE user SET password = ? WHERE uid = ?")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(password, uid)
	if err != nil {
		return fail.NewMySqlError(err)
	}

	return nil
}

func (u *userRepository) AlterNick(nick string, uid []byte) *fail.ResponseError {
	defer u.db.Close()

	stmt, err := u.db.Prepare("UPDATE user SET nick = ? WHERE uid = ?")
	if err != nil {
		return fail.NewMySqlError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(nick, uid)
	if err != nil {
		return fail.NewMySqlError(err)
	}

	return nil
}

func (u *userRepository) SearchUsers() (*[]user, *fail.ResponseError) {

	defer u.db.Close()

	rows, err := u.db.Query("SELECT first_name, nick FROM user")
	if err != nil {
		return nil, fail.NewMySqlError(err)
	}

	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user

		err = rows.Scan(
			&u.First_name,
			&u.Nick,
		)
		if err != nil {
			fmt.Println(err)
			return nil, fail.NewMySqlError(err)
		}

		users = append(users, u)
	}

	return &users, nil
}

func (u *userRepository) SearchUserByID(UID []byte) (*entities.User, *fail.ResponseError) {
	defer u.db.Close()

	row := u.db.QueryRow("SELECT first_name, nick FROM user WHERE uid = ?", UID)

	var user entities.User

	err := row.Scan(&user)
	if err == sql.ErrNoRows {
		return nil, &fail.ResponseError{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
		}
	}

	if err != nil {
		return nil, fail.NewMySqlError(err)
	}

	return &user, nil
}
