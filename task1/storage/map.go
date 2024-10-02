package storage

import (
	. "GoHomework/task1/book"
)

type MapStorage struct {
	data map[uint64]Book
}

func NewMapStorage() *MapStorage {
	return &MapStorage{data: make(map[uint64]Book)}
}

func (storage *MapStorage) Size() int {
	return len(storage.data)
}

func (storage *MapStorage) Add(book Book) {
	id := book.ID()
	storage.data[id] = book
}

func (storage *MapStorage) Get(id uint64) (Book, bool) {
	book, ok := storage.data[id]
	return book, ok
}

func (storage *MapStorage) Remove(id uint64) bool {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return true
	}
	return false
}

func (storage *MapStorage) GetData() []Book {
	data := make([]Book, 0, len(storage.data))
	for _, book := range storage.data {
		data = append(data, book)
	}
	return data
}
func (storage *MapStorage) Clear() {
	storage.data = make(map[uint64]Book)
}
