package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/cors"
	"time"
	globals "gin_session_auth/globals"
	middleware "gin_session_auth/middleware"
	routes "gin_session_auth/routes"
)



func main() {
	router := gin.Default()
	config := cors.Config {
		AllowOrigins: []string {
			"http://127.0.0.1:3333",
			"http://127.0.0.1:3000",
		},
		AllowMethods: []string {
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowOriginFunc: func(origin string) bool {
      return origin == "http://127.0.0.1:3000"
    },
		AllowCredentials: true,
		MaxAge: 1 * time.Hour,
	}
	router.Use(cors.New(config))
	router.SetTrustedProxies([]string{"localhost:3333"})
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)
	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	

	router.Run(":3333")
}