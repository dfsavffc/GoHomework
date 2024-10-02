package book

type Book struct {
	Title   string
	Content string
	Length  int
}

func NewBook(title, content string, length int) *Book {
	return &Book{
		Title:   title,
		Content: content,
		Length:  length,
	}
}
