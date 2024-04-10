package routes

import "net/http"

var PostsRoutes = [...]Route{
	{
		URI:            http.MethodGet + "/posts",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: false,
	},
	{
		URI:            http.MethodGet + "/posts/{post_id}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/posts",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/posts/{post_id}/comment",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPatch + "/posts/{post_id}/comment",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodDelete + "/posts/{post_id}/comment",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/posts/{post_id}/react",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/posts/{post_id}/share",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/posts/{post_id}/report",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPatch + "/posts/{post_id}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodDelete + "/posts/{post_id}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
}
