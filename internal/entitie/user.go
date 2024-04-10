package entitie

import "time"

type User struct {
	UID       string   `json:"uid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birthDate"`
}
