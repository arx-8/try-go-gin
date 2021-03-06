package handlers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/arx-8/try-go-gin/src/handlers/requests"
	"github.com/arx-8/try-go-gin/src/model"
	"github.com/arx-8/try-go-gin/src/service"
	"github.com/gin-gonic/gin"
)

type BooksHandlerInterface interface {
	GetList(context *gin.Context)
	GetByID(context *gin.Context)
	PostNew(context *gin.Context)
	Delete(context *gin.Context)
	Put(context *gin.Context)
}

type BooksHandler struct {
	bookService service.BookService
}

func NewBooksHandler(bookService service.BookService) BooksHandlerInterface {
	return &BooksHandler{bookService: bookService}
}

func (h *BooksHandler) GetList(c *gin.Context) {
	// get query params
	var getBooks requests.GetBooks

	if err := c.ShouldBindQuery(&getBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		fmt.Println(getBooks)
		fmt.Println(reflect.TypeOf(getBooks.Start))
	}

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

func (h *BooksHandler) PostNew(c *gin.Context) {
	var req requests.PostNewBook
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := h.bookService.Add(req.ToPlain())

	c.JSON(http.StatusCreated, gin.H{
		"id": id.ToInt(),
	})
}

func (h *BooksHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := model.BookIDFromString(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID:" + idParam,
		})
		return
	}

	if err := h.bookService.DeleteByID(id); err != nil {
		// TODO ??????????????? NotFound ?????????????????????????????????????????? errors.Is, errors.As ?
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BooksHandler) Put(c *gin.Context) {
	idParam := c.Param("id")
	id, err := model.BookIDFromString(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID:" + idParam,
		})
		return
	}

	var req requests.PostNewBook
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.bookService.UpdateByID(id, req.ToPlain()); err != nil {
		// TODO ??????????????? NotFound ?????????????????????????????????????????? errors.Is, errors.As ?
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
