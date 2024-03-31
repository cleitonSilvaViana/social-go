// package routes contains the routes of application configurated
package routes

import (
	"net/http"
)


var UsersRoutes = [...]Route{
	{
		URI:            http.MethodPost + "/users",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: false,
	},
	{
		URI:            http.MethodGet + "/users",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: false,
	},
	{
		URI:            http.MethodGet + "/users/{user-uid}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: false,
	},
	{
		URI:            http.MethodPost + "/{user-uid}/followers",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/{user-id}/following-me",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},

	{
		URI:            http.MethodPost + "/users/{user-uid}/invite",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
		{
		URI:            http.MethodPost + "/users/{user-uid}/block",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPatch + "/users/{user-id}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/users/{user-uid}/update-password",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/users/{user-uid}/disable",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodDelete + "/users/{user-uid}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
}