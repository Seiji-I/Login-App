package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"

	globals "gin_session_auth/globals"
	// middleware "gin_session_auth/middleware"
	routes "gin_session_auth/routes"
)



func main() {
	router := gin.Default()

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	router.Run(":3333")
}