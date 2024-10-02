package storage

import (
	. "GoHomework/task1/book"
)

type ModelStorage interface {
	Add(Book)
	Remove(uint64) bool
	Get(uint64) (Book, bool)
	Size() int
	GetData() []Book
}
