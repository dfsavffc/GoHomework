package storage

import (
	"GoHomework/task1/book"
)

type ModelStorage interface {
	Add(book.Book)
	Remove(uint64) error
	Get(uint64) (book.Book, error)
	Size() int
	GetData() []book.Book
	Clear()
}
