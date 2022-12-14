package routes

import (
	"github.com/gin-gonic/gin"
	controllers "gin_session_auth/controllers"

)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/logout", controllers.LoginGetHandler())
	g.GET("/", controllers.IndexGetHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardGetHandler())
}