package entitie

import "time"

type Profile struct {
	UID       string    `json:"uid"`
	ImgPerfil []byte    `json:"imgPerfil"`
	Nick      string    `json:"nick"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	AddressID int       `json:"AddressId"`
	ContactID int       `json:"contactId"`
}

/*

type Profile struct {
 	UID []byte
 	Nick string
 	FirstName string
 	LastName string
 	Password []byte
	CreatedAt time.Time
	Email string
}

*/
