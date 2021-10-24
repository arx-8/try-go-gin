package main

import (
	"github.com/arx-8/try-go-gin/src/handler/books"
	"github.com/arx-8/try-go-gin/src/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.RecordUaAndTime)

	booksRoute := r.Group("books")

	booksRoute.GET("", books.GetList)
	booksRoute.POST("")
	booksRoute.GET("/:id", books.GetByID)
	booksRoute.PUT("/:id")
	booksRoute.DELETE("/:id")

	r.GET("/healthz", func(c *gin.Context) {
		c.Done()
	})

	r.Run(":8080")
}
