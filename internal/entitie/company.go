package entitie

import "time"

type Companym struct {
	UID        string    `json:"uid"`
	Foundation time.Time `json:"foundation"`
}
