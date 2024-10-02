package main

import (
	. "GoHomework/book"
	"GoHomework/library"
	"GoHomework/storage"
	"math"
)

func LenGenerator(string string) int {
	return len(string)
}
func SquareLenGenerator(string string) int {
	return int(math.Pow(float64(len(string)), 10))
}
func main() {
	aBook := Book{Title: "aa", Content: "aa", Length: 2}
	bBook := Book{Title: "bbbbb", Content: "bbbbb", Length: 8}
	cBook := Book{Title: "ccccccc", Content: "ccccccccc", Length: 13}
	dBook := Book{Title: "dddddddddd", Content: "ddddddd", Length: 14}
	eBook := Book{Title: "eeeeeeeeeeeeee", Content: "eeeeeeeeeeeeee", Length: 15}

	mapStorage := storage.NewMapStorage()
	sliceStorage := storage.NewSliceStorage()

	lib := library.NewLibrary(mapStorage, LenGenerator)
	lib.SetIdGenerator(LenGenerator)

	lib.Add(aBook)
	lib.Add(bBook)
	println(lib.Size())
	lib.Add(cBook)
	lib.Add(dBook)

	book, ok := lib.Get("bbbbb")
	println(ok, book.Content, book.Length)

	book, ok = lib.Get("dddddddddd")
	println(ok, book.Content, book.Length)

	lib.Remove("dddddddddd")

	book, ok = lib.Get("dddddddddd")
	println(ok)

	lib.Add(dBook)

	lib.SetIdGenerator(SquareLenGenerator)
	lib.Add(eBook)

	book, ok = lib.Get("bbbbb")
	println(ok, book.Content, book.Length)

	book, ok = lib.Get("eeeeeeeeeeeeee")
	println(ok, book.Content, book.Length)

	lib = library.NewLibrary(sliceStorage, LenGenerator)
	println(lib.Size())

	lib.Add(aBook)
	lib.Add(bBook)
	lib.Add(cBook)
	lib.Add(dBook)

	book, ok = lib.Get("bbbbb")
	println(ok, book.Content, book.Length)

	book, ok = lib.Get("dddddddddd")
	println(ok, book.Content, book.Length)
}
