package storage

import (
	"GoHomework/task1/book"
	"errors"
	"fmt"
)

type SliceStorage struct {
	data []book.Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{data: make([]book.Book, 0)}
}

func (storage *SliceStorage) Add(book book.Book) {
	storage.data = append(storage.data, book)
}

func (storage *SliceStorage) Get(id uint64) (book.Book, error) {
	for _, elem := range storage.data {
		if elem.ID() == id {
			return elem, nil
		}
	}
	return book.Book{}, errors.New(fmt.Sprintf("id:%d not exist", id))
}

func (storage *SliceStorage) Remove(id uint64) error {
	for i, elem := range storage.data {
		if elem.ID() == id {
			storage.data = append(storage.data[:i], storage.data[i+1:]...)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("id:%d not exist", id))
}

func (storage *SliceStorage) Size() int {
	return len(storage.data)
}

func (storage *SliceStorage) GetData() []book.Book {
	return storage.data
}

func (storage *SliceStorage) Clear() {
	storage.data = make([]book.Book, 0)
}
