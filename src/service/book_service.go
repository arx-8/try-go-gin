package service

import "github.com/arx-8/try-go-gin/src/model"

var books = []model.Book{
	{
		Id:      0,
		Title:   "Full Metal Panic",
		Content: "Robot science fiction",
	},
	{
		Id:      1,
		Title:   "Bitcoin road",
		Content: "The bitcoin road will be long",
	},
	{
		Id:      2,
		Title:   "Black swans",
		Content: "All markets are randomness",
	},
}

type BookService struct{}

func (BookService) GetList() []model.Book {
	return books
}
