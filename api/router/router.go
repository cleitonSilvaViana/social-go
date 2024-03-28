// package router contains the router on the application
package router

import (
	"github.com/cleitonSilvaViana/social-go/api/router/routes"
	"github.com/cleitonSilvaViana/social-go/internal/config"
	"github.com/gin-gonic/gin"
)

// StartRouter init the process of routing in the application
func StartRouter(APIPost string) {
	err := config.GetConfig()
	if err != nil {
		panic(err)
	}


	var router *gin.Engine = gin.Default()
	routes.Users(router)
	router.Run(APIPost)
}
