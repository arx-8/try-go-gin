package main

import (
	"github.com/arx-8/try-go-gin/src/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.RecordUaAndTime)

	booksRoute := r.Group("books")

	booksRoute.GET("")
	booksRoute.POST("")
	booksRoute.GET("/:id")
	booksRoute.PUT("/:id")
	booksRoute.DELETE("/:id")

	r.GET("/healthz", func(c *gin.Context) {
		c.Done()
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// You can access `curl http://localhost:8080/healthz`
}
