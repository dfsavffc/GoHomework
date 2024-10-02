package storage

import (
	. "GoHomework/task1/book"
)

type SliceStorage struct {
	data []Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{data: make([]Book, 0)}
}

func (storage *SliceStorage) Add(book Book) {
	storage.data = append(storage.data, book)
}

func (storage *SliceStorage) Get(id uint64) (Book, bool) {
	for _, book := range storage.data {
		if book.ID() == id {
			return book, true
		}
	}
	return Book{}, false
}

func (storage *SliceStorage) Remove(id uint64) bool {
	for i, book := range storage.data {
		if book.ID() == id {
			storage.data = append(storage.data[:i], storage.data[i+1:]...)
			return true
		}
	}
	return false
}

func (storage *SliceStorage) Size() int {
	return len(storage.data)
}

func (storage *SliceStorage) GetData() []Book {
	return storage.data
}
