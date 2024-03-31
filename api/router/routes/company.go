package routes

import "net/http"

var CompanyRoutes = [...]Route{
	{
		URI:            http.MethodGet + "/company",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: false,
	},
	{
		URI:            http.MethodGet + "/company/{company-uid}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPost + "/company",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodPatch + "/company/{company-uid}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
	{
		URI:            http.MethodDelete + "/company/{company-uid}",
		Handler:        func(w http.ResponseWriter, r *http.Request) {},
		AuthIsrequired: true,
	},
}
