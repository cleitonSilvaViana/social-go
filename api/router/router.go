// package router contains the router on the application
package router

import "github.com/gin-gonic/gin"

// StartRouter init the process of routing in the application
func StartRouter() {
	var router *gin.Engine = gin.Default()
	router.Run(":5000")
}
