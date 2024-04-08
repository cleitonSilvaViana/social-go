package handlers

import (
	"io"
	"net/http"

	"github.com/cleitonSilvaViana/social-go/internal/domain"
	"github.com/cleitonSilvaViana/social-go/internal/repository"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, e := io.ReadAll(r.Body)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e.Error()))
		return
	}

	err := domain.CreateUser(body)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := domain.SearchUsers()
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	userUID := r.PathValue("user_id")

	user, err := domain.SearchUser(userUID)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")

	userUID := r.PathValue("user_id")

	// realizar as validações aqui...

	repo, err := repository.NewUserRepository()
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	err = repo.DeleteUser(userUID)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
