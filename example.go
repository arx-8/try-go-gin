package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ua := ""

	r.Use(func(c *gin.Context) {
		ua = c.GetHeader("user-agent")
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "pong",
			"user-agent": ua,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// You can access `http://localhost:8080/ping`
}
