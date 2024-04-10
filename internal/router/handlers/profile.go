package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cleitonSilvaViana/social-go/internal/domain"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, e := io.ReadAll(r.Body)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e.Error()))
		return
	}

	token, err := domain.Login(body)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func Logout(w http.ResponseWriter, r *http.Request) {}

func Report(w http.ResponseWriter, r *http.Request) {}

func Follow(w http.ResponseWriter, r *http.Request) {}

func Unfollow(w http.ResponseWriter, r *http.Request) {}

func SearchProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ler query params aqui
	var queryParams map[string]string

	parsedQuery, e := url.ParseQuery(r.URL.Path)
	if e != nil {
		// gerar log aqui...

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return 
	}

	for key, values := range parsedQuery {
		for _, value := range values {
			queryParams[key] = value
		}
	}


	users, err := domain.SearchProfiles(queryParams)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func SearchProfile(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	userUID := r.PathValue("profile_id")

	param := map[string]string{
		"uid": userUID,
	}

	user, err := domain.SearchProfile(param)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Message))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user))
}
