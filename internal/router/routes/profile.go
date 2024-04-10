package routes

import (
	"net/http"

	"github.com/cleitonSilvaViana/social-go/internal/router/handlers"
)

type Route struct {
	URI            string
	Handler        func(w http.ResponseWriter, r *http.Request)
	AuthIsrequired bool
}

// ProfileRoutes contains the routes in common between users and companies
var ProfileRoutes = [...]Route{
	{
		URI:            http.MethodPost + "/login",
		Handler:        handlers.Login,
		AuthIsrequired: false,
	},
	{
		URI:            http.MethodPost + "/logout",
		Handler:        handlers.Logout,
		AuthIsrequired: true,
	},
	{
		URI:            "GET /profiles",
		Handler:        handlers.SearchProfiles,
		AuthIsrequired: false,
	},
	{
		URI:            "GET /profiles/{profile_id}",
		Handler:        handlers.SearchProfile,
		AuthIsrequired: false,
	},
	{
		URI:            http.MethodPost + "/profiles/report/{profile_id}",
		Handler:        handlers.Report,
		AuthIsrequired: true,
	},

	{
		URI:            http.MethodPost + "profiles/follow/{profile_id}",
		Handler:        handlers.Follow,
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "profiles/unfollow/{profile_id}",
		Handler:        handlers.Unfollow,
		AuthIsrequired: true,
	},
	{
		URI:            "POST /profiles/{profile_id}/disable",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            "PATCH /profiles/{profile_id}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            "POST /profiles/{profile_id}/update/password",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            "DELETE /profiles/{profile_id}",
		Handler:        handlers.DeleteUser,
		AuthIsrequired: true,
	},
	{
		URI:            "POST /profiles/block/{profile_id}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            "POST /profiles/{user_id}/invite",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
}
