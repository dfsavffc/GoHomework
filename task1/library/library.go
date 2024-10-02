package library

import (
	. "GoHomework/task1/book"
	"GoHomework/task1/identification"
	"GoHomework/task1/storage"
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
	for _, book := range data {
		lib.Add(book)
	}
}

func (lib *Library) SetIdGenerator(generator identification.Generator) {
	lib.IdGenerator = generator
}

func (lib *Library) Size() int {
	return lib.storage.Size()
}

func (lib *Library) Get(title string) (Book, bool) {
	for _, book := range lib.storage.GetData() {
		if book.Title == title {
			return book, true
		}
	}
	return Book{}, false
}

func (lib *Library) Add(book Book) {
	id := lib.IdGenerator()
	book.SetID(id)
	lib.storage.Add(book)
}

func (lib *Library) Remove(id uint64) bool {
	if lib.storage.Remove(id) {
		return true
	}
	return false
}
