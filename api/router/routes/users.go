// package routes contains the routes of application configurated
package routes

import "github.com/gin-gonic/gin"

// User contains the users routes
func Users(router *gin.Engine) {

	// notRequired contains the routes acessible for any user
	var notRequired *gin.RouterGroup = router.Group("/users")

	notRequired.POST("", func(ctx *gin.Context) {}) // route responsibility for create a new user
	notRequired.POST("/login", func(ctx *gin.Context) {})

	// required contains the routes wath the user authenticate is necesary
	var required *gin.RouterGroup = router.Group("/users")

	required.GET("", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("hiiiiiiiiiiiiiiiiiiii"))
	}) // search users
	required.GET("/:id", func(ctx *gin.Context) {}) // search one user
	required.PATCH("/:uid", func(ctx *gin.Context) {})
	required.DELETE("/:uid", func(ctx *gin.Context) {})
}
