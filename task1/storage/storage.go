package storage

import (
	. "GoHomework/task1/book"
)

type IdGenerator func(string) int

type ModelStorage interface {
	Add(Book, int)
	Remove(int) bool
	Get(int) (Book, bool)
	Size() int
	GetAll() []Book
	Clear()
}
