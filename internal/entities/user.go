package entities

import "time"

type User struct {
	UID       string   `json:"uid"`
	Nick      string    `json:"nick"`
	ImgPerfil []byte    `json:"imgPerfil"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	ContactID int       `json:"contactId"`
	CityID    int       `json:"cityId"`
	Password  []byte    `json:"password"`
}
