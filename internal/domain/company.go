package domain

import (
	"time"

	"github.com/cleitonSilvaViana/social-go/pkg/fail"
)

type Company struct {
	Foundation time.Time `json:"foundation"`
}

func CreateCompany() *fail.ResponseError {

	return nil
}

