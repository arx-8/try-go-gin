package handlers

import (
	"net/http"

	"github.com/arx-8/try-go-gin/src/model"
	"github.com/arx-8/try-go-gin/src/service"
	"github.com/gin-gonic/gin"
)

type BooksHandlerInterface interface {
	GetList(context *gin.Context)
	GetByID(context *gin.Context)
}

type BooksHandler struct {
	bookService service.BookService
}

func NewBooksHandler(bookService service.BookService) BooksHandlerInterface {
	return &BooksHandler{bookService: bookService}
}

func (h *BooksHandler) GetList(c *gin.Context) {
	books := h.bookService.GetList()

	c.JSON(http.StatusOK, gin.H{
		"data":               books,
		"total_title_length": model.GetTotalTitleLength(books),
	})
}

func (h *BooksHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := model.BookIDFromString(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID:" + idParam,
		})
		return
	}

	book, err := h.bookService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
