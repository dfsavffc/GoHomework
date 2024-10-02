package book

type Book struct {
	Title   string
	Author  string
	Content string
	id      *uint64
}

func NewBook(title, content string, author string) *Book {
	return &Book{
		Title:   title,
		Content: content,
		Author:  author,
		id:      nil,
	}
}

func (book *Book) ID() uint64 {
	return *book.id
}

func (book *Book) SetID(id uint64) {
	book.id = &id
}
