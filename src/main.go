package main

import (
	"github.com/arx-8/try-go-gin/src/handlers"
	"github.com/arx-8/try-go-gin/src/middleware"
	"github.com/arx-8/try-go-gin/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.RecordUaAndTime)

	{
		booksRoute := r.Group("books")
		booksHandler := handlers.NewBooksHandler(service.BookService{})
		booksRoute.GET("", booksHandler.GetList)
		booksRoute.POST("", booksHandler.PostNew)
		booksRoute.GET("/:id", booksHandler.GetByID)
		booksRoute.PUT("/:id")
		booksRoute.DELETE("/:id")
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.Done()
	})

	r.Run(":8080")
}
