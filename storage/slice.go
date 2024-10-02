package storage

import (
	. "GoHomework/book"
)

type SliceStorage struct {
	data []Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{data: make([]Book, 0)}
}

func (storage *SliceStorage) Add(book Book, id int) {
	if id >= len(storage.data) {
		newData := make([]Book, id+1)
		copy(newData, storage.data)
		storage.data = newData
	}
	storage.data[id] = book
}

func (storage *SliceStorage) Remove(id int) bool {
	if id < len(storage.data) {
		storage.data[id] = Book{} // Заменяем на пустую книгу
		return true
	}
	return false
}

func (storage *SliceStorage) Get(id int) (Book, bool) {
	if id < len(storage.data) {
		return storage.data[id], true
	}
	return Book{}, false
}

func (storage *SliceStorage) Size() int {
	return len(storage.data)
}

func (storage *SliceStorage) GetAll() []Book {
	return storage.data
}

func (storage *SliceStorage) Clear() {
	storage.data = make([]Book, 0)
}
