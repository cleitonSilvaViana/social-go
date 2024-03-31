// package router contains the router on the application
package router

import (
	"log"
	"net/http"

	"github.com/cleitonSilvaViana/social-go/api/router/routes"
)

func InitRouter(apiPort string) {
	mux := http.NewServeMux()
	for _, route := range routes.UsersRoutes {
		mux.HandleFunc(route.URI, route.Handler)
	}
	log.Fatal(http.ListenAndServe(apiPort, mux))
}
