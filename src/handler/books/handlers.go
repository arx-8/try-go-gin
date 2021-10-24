package books

import (
	"net/http"

	"github.com/arx-8/try-go-gin/src/model"
	"github.com/arx-8/try-go-gin/src/service"
	"github.com/gin-gonic/gin"
)

var bookService = service.BookService{}

func GetList(c *gin.Context) {
	books := bookService.GetList()

	c.JSON(http.StatusOK, gin.H{
		"data":               books,
		"total_title_length": model.GetTotalTitleLength(books),
	})
}

func GetByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := model.BookIDFromString(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID:" + idParam,
		})
		return
	}

	book, err := bookService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
