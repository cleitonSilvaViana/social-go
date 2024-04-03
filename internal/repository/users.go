package repository

import (
	"database/sql"
	"fmt"

	"github.com/cleitonSilvaViana/social-go/internal/database"
)

type user struct {
	first_name string
	last_name  string
	nick       string
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() (*userRepository, error) {
	db, err := database.ConnectToMysql()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &userRepository{db}, nil
}

func (u *userRepository) CreateUser() error {

	defer u.db.Close()

	return nil
}

func (u *userRepository) SearchUsers() (*[]user, error) {

	defer u.db.Close()

	rows, err := u.db.Query("SELECT first_name, last_name, nick FROM user")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user

		err = rows.Scan(
			&u.first_name,
			&u.last_name,
			&u.nick,
		)
		if err != nil {
			fmt.Println("SEARCHUSERS ERRO 2 ")
			return nil, err
		}

		users = append(users, u)
	}

	return &users, nil
}
