// package routes contains the routes of application configurated
package routes

import (
	"net/http"

	"github.com/cleitonSilvaViana/social-go/api/handlers"
)


var UsersRoutes = [...]Route{
	{
		URI:            "POST /users",
		Handler:        handlers.RegisterUser,
		AuthIsrequired: false,
	},
	{
		URI:            "GET /users",
		Handler:        handlers.SearchUsers,
		AuthIsrequired: false,
	},
	{
		URI:            "GET /users/{user_id}",
		Handler:        handlers.SearchUser,
		AuthIsrequired: false,
	},
	{
		URI:           "POST /users/{user_id}/followers",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:           "POST /users/{user_id}/following-me",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},

	{
		URI:           "POST /users/{user_id}/invite",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
		{
		URI:           "POST /users/{user_id}/block",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            "PATCH /users/{user_id}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:           "POST /users/{user_id}/update-password",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:           "POST /users/{user_id}/disable",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            "DELETE /users/{user_id}",
		Handler:        handlers.DeleteUser,
		AuthIsrequired: true,
	},
}