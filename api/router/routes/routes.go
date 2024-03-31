package routes

import "net/http"

type Route struct {
	URI            string
	Handler        func(w http.ResponseWriter, r *http.Request)
	AuthIsrequired bool
}

// ProfileRoutes contains the routes in common between users and companies
var ProfileRoutes = [...]Route{
	{
		URI:            http.MethodPost + "/login",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/logout",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/report/profile-uid",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},

	{
		URI:            http.MethodPost + "/follow/profile-uid",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/unfollow/profile-uid",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
}



