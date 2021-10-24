package service

import (
	"github.com/arx-8/try-go-gin/src/model"
)

var books = []model.Book{
	{
		ID:      0,
		Title:   "Full Metal Panic",
		Content: "Robot science fiction",
	},
	{
		ID:      1,
		Title:   "Bitcoin road",
		Content: "The bitcoin road will be long",
	},
	{
		ID:      2,
		Title:   "Black swans",
		Content: "All markets are randomness",
	},
}

type BookService struct{}

func (BookService) GetList() []model.Book {
	return books
}

func (BookService) GetByID(id model.BookID) (*model.Book, error) {
	for _, v := range books {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, ErrRecordNotFound("ID:" + id.ToString() + " is not found.")
}

func (BookService) Add(data struct {
	Title   string
	Content string
}) model.BookID {
	nextID := model.BookID(books[len(books)-1].ID.ToInt() + 1)

	books = append(
		books,
		*model.NewBook(nextID, data.Title, data.Content),
	)

	return nextID
}
