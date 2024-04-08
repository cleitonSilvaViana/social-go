// package router contains the router on the application
package router

import (
	"log"
	"net/http"

	middleware "github.com/cleitonSilvaViana/social-go/api/middlewares"
	"github.com/cleitonSilvaViana/social-go/api/router/routes"
)

// InitRouter realize the routing stater in application
func InitRouter(apiPort string) {
	mux := http.NewServeMux()
	for _, route := range routes.UsersRoutes {
		mux.HandleFunc(route.URI, middleware.LoggingMiddleware(route.Handler))
	}

	log.Fatal(http.ListenAndServe(apiPort, mux))
}
