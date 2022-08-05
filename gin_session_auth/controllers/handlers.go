package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	globals "gin_session_auth/globals"
	helpers "gin_session_auth/helpers"
)

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println(user)
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
		c.Bind(&username)
		c.Bind(&password)
		log.Printf("%s\n", username)
		log.Printf("%s\n", password)
		if helpers.EmptyUsePass(username, password) {
			c.JSON(http.StatusBadRequest, gin.H {
				"content": "Parameters can't be empty",
			})
			log.Println("Parameters can't be empty")
			return
		}

		if !helpers.CheckUserPass(username, password) {
			c.JSON(http.StatusInternalServerError, gin.H {
				"content": "Incorrect username or password",
			})
			log.Println("Incorrect username or password")
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"content": "Failed to save session",
			})
			log.Println("Failed to save session")
			return
		}
		c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:3000")
			return
		}
		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:3000")
	}
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.JSON(http.StatusOK, gin.H{
			"content": "This is an index page...",
			"user": user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.JSON(http.StatusOK, gin.H {
			"content": "This is a dashboard",
			"user": user,
		})
	}
}