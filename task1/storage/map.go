package storage

import (
	"GoHomework/task1/book"
	"errors"
	"fmt"
)

type MapStorage struct {
	data map[uint64]book.Book
}

func NewMapStorage() *MapStorage {
	return &MapStorage{data: make(map[uint64]book.Book)}
}

func (storage *MapStorage) Size() int {
	return len(storage.data)
}

func (storage *MapStorage) Add(book book.Book) {
	id := book.ID()
	storage.data[id] = book
}

func (storage *MapStorage) Get(id uint64) (book.Book, error) {
	elem, ok := storage.data[id]
	if !ok {
		return book.Book{}, errors.New(fmt.Sprintf("id:%d not exist", id))
	}
	return elem, nil
}

func (storage *MapStorage) Remove(id uint64) error {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return nil
	}
	return errors.New(fmt.Sprintf("id:%d not exist", id))
}

func (storage *MapStorage) GetData() []book.Book {
	data := make([]book.Book, 0, len(storage.data))
	for _, elem := range storage.data {
		data = append(data, elem)
	}
	return data
}
func (storage *MapStorage) Clear() {
	storage.data = make(map[uint64]book.Book)
}
