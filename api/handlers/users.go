package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cleitonSilvaViana/social-go/internal/domain"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print("ERRO 1")
		// tratar error - bad request
		return
	}

	err = domain.CreateUser(body)
	if err != nil {
		fmt.Print("ERRO 2")
		// tratar erro
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

 	users, err := domain.SearchUsers()
	if err != nil {
		fmt.Println("ERRO 3")
		// tratar error
		return
	} 

	w.WriteHeader(http.StatusOK)
	w.Write(users)
}
