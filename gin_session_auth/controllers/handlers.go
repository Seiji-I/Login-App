package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"

	globals "gin_session_auth/globals"
	helpers "gin_session_auth/helpers"
)

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"content": "please logout first",
				"user": user,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"content": "",
			"user": user,
		})
	}
}


func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUsePass(username, password) {
			c.JSON(http.StatusBadRequest, gin.H {
				"content": "Parameters can't be empty",
			})
			return
		}

		if !helpers.CheckUserPass(username, password) {
			c.JSON(http.StatusInternalServerError, gin.H {
				"content": "Incorrect username or password",
			})
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": "Failed to save session",
			})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "localhost:3000/home")
	}
}