// package routes contains the routes of application configurated
package routes

import (
	"net/http"

	"github.com/cleitonSilvaViana/social-go/internal/router/handlers"
)

var UsersRoutes = [...]Route{
	{
		URI:            "POST /users",
		Handler:        handlers.RegisterUser,
		AuthIsrequired: false,
	},
	{
		URI:            "POST /users/{user_id}/followers",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            "POST /users/{user_id}/following-me",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},

}
