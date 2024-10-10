package library

import (
	"GoHomework/task1/book"
	"GoHomework/task1/identification"
	"GoHomework/task1/storage"
	"errors"
	"fmt"
)

type Library struct {
	storage     storage.ModelStorage
	IdGenerator identification.Generator
}

func NewLibrary(storage storage.ModelStorage) *Library {
	return &Library{storage: storage, IdGenerator: identification.DefaultGenerator()}
}

func (lib *Library) ReIdentification() {
	data := lib.storage.GetData()
	lib.storage.Clear()
	for _, elem := range data {
		lib.Add(elem)
	}
}

func (lib *Library) SetIdGenerator(generator identification.Generator) {
	lib.IdGenerator = generator
}

func (lib *Library) Size() int {
	return lib.storage.Size()
}

func (lib *Library) Get(title string) (book.Book, error) {
	for _, elem := range lib.storage.GetData() {
		if elem.Title == title {
			return elem, nil
		}
	}
	return book.Book{}, errors.New(fmt.Sprintf("title:%s not exist", title))
}

func (lib *Library) Add(book book.Book) {
	id := lib.IdGenerator()
	book.SetID(id)
	lib.storage.Add(book)
}

func (lib *Library) Remove(id uint64) error {
	ok := lib.storage.Remove(id)
	return ok
}
