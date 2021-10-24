package model

import (
	"strconv"
)

type BookID int

func BookIDFromString(maybeID string) (BookID, error) {
	idAsInt, err := strconv.Atoi(maybeID)
	if err != nil {
		return -1, err
	}

	return BookID(idAsInt), nil
}

func (id BookID) ToString() string {
	return strconv.Itoa(id.ToInt())
}

func (id BookID) ToInt() int {
	return int(id)
}

type Book struct {
	ID      BookID `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewBook(id BookID, title string, content string) *Book {
	n := new(Book)
	n.ID = id
	n.Title = title
	n.Content = content
	return n
}

func (b Book) GetTitleLength() int {
	return len(b.Title)
}

func GetTotalTitleLength(books []Book) int {
	result := 0
	for _, v := range books {
		result += v.GetTitleLength()
	}
	return result
}
