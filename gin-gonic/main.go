package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	id string
	name string
	password string
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})
	router.GET("/dbtest", func(c *gin.Context) {
		db, err := sql.Open("postgres-db", "host=127.0.0.1 port=5432 user=root password=root dbname=testdb")
		defer db.Close()

		if err != nil {
			fmt.Println("OK")
			c.JSON(200, gin.H{
				"message": "connect to postgresql was success",
			})
		} else {
			c.JSON(200, gin.H{
				"message": err,
			})
		}
	})

	router.Run(":3333")
}