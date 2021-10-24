package requests

type PostNewBook struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

func (p PostNewBook) ToPlain() struct {
	Title   string
	Content string
} {
	return struct {
		Title   string
		Content string
	}{
		Title:   p.Title,
		Content: p.Content,
	}
}
