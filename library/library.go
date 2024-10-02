package library

import (
	. "GoHomework/book"
	"GoHomework/storage"
)

type ReLoader interface {
	Reload()
}

type ModelLib interface {
	Size() int
	Get(string) (Book, bool)
	Add(book Book)
	Remove(string) bool
	Id(string) int
}
type Library struct {
	storage     storage.ModelStorage
	IdGenerator storage.IdGenerator
}

func (lib *Library) SetIdGenerator(generator storage.IdGenerator) {
	lib.IdGenerator = generator
	lib.Reload()
}

func NewLibrary(storage storage.ModelStorage, generator storage.IdGenerator) *Library {
	return &Library{storage: storage, IdGenerator: generator}
}

func (lib *Library) Reload() {
	data := lib.storage.GetAll()
	lib.storage.Clear()
	for _, book := range data {
		id := lib.IdGenerator(book.Title)
		lib.storage.Add(book, id)
	}
}

func (lib *Library) Size() int {
	return lib.storage.Size()
}
func (lib *Library) Get(title string) (Book, bool) {
	id := lib.IdGenerator(title)
	return lib.storage.Get(id)
}
func (lib *Library) Add(book Book) {
	id := lib.IdGenerator(book.Title)
	lib.storage.Add(book, id)
}
func (lib *Library) Remove(title string) bool {
	id := lib.IdGenerator(title)
	return lib.storage.Remove(id)
}
