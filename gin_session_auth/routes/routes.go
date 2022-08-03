package routes

import (
	"github.com/gin-gonic/gin"
	controllers "gin_session_auth/controllers"

)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
}