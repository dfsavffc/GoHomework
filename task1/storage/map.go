package storage

import (
	. "GoHomework/task1/book"
)

type MapStorage struct {
	data map[int]Book
}

func NewMapStorage() *MapStorage {
	return &MapStorage{data: make(map[int]Book)}
}

func (storage *MapStorage) Clear() {
	storage.data = make(map[int]Book)
}

func (storage *MapStorage) GetAll() []Book {
	data := make([]Book, 0, len(storage.data))
	for _, b := range storage.data {
		data = append(data, b)
	}
	return data
}

func (storage *MapStorage) Size() int {
	return len(storage.data)
}

func (storage *MapStorage) Add(book Book, id int) {
	storage.data[id] = book
}

func (storage *MapStorage) Remove(id int) bool {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return true
	}
	return false
}

func (storage *MapStorage) Get(id int) (Book, bool) {
	b, ok := storage.data[id]
	return b, ok
}
