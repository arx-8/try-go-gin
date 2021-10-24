package model

type Book struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewBook(id int, title string, content string) *Book {
	n := new(Book)
	n.Id = id
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
