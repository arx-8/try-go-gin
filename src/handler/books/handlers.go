package books

import (
	"net/http"

	"github.com/arx-8/try-go-gin/src/model"
	"github.com/arx-8/try-go-gin/src/service"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	bookService := service.BookService{}
	books := bookService.GetList()

	c.JSON(http.StatusOK, gin.H{
		"data":               books,
		"total_title_length": model.GetTotalTitleLength(books),
	})
}
