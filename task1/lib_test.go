package main

import (
	. "GoHomework/task1/book"
	"GoHomework/task1/identification"
	"GoHomework/task1/library"
	"GoHomework/task1/storage"
	"testing"
)

func TestMapLibrary(t *testing.T) {
	mapStorage := storage.NewMapStorage()
	lib := library.NewLibrary(mapStorage)

	justBook := Book{Title: "Just Book"}
	simpleBook := Book{Title: "Simple Book"}
	normalBook := Book{Title: "Normal Book"}
	hardBook := Book{Title: "Hard Book"}
	specialBook := Book{Title: "Special Book"}
	expensiveBook := Book{Title: "Expensive Book", Content: "?"}
	goodBook := Book{Title: "Good Book"}
	wonderfulBook := Book{Title: "Wonderful Book"}

	lib.Add(justBook)
	lib.Add(simpleBook)
	lib.Add(normalBook)
	lib.Add(wonderfulBook)
	lib.Add(expensiveBook)
	lib.Add(goodBook)

	if lib.Size() != 6 {
		t.Errorf("Lib.Size() = %d; excepted 6", lib.Size())
	}

	book, _ := lib.Get("Simple Book")
	if book.Title != "Simple Book" {
		t.Errorf("book.Title = %s; excepted 'Simple Book' ", book.Title)
	}

	book, _ = lib.Get("Expensive Book")
	if book.Content != "?" {
		t.Errorf("book.Content = %s; excepted '?' ", book.Content)
	}

	book, _ = lib.Get("Expensive Book")
	id := book.ID()
	lib.Remove(id)

	if lib.Size() != 5 {
		t.Errorf("Lib.Size() = %d; excepted 5", lib.Size())
	}

	book, ok := lib.Get("Expensive Book")
	if ok {
		t.Errorf("book.Title = %s, ok = true; excepted '', ok = false ", book.Title)
	}

	lib.SetIdGenerator(identification.AlmostDefaultGenerator())
	lib.Add(hardBook)
	lib.Add(specialBook)

	book, _ = lib.Get("Hard Book")
	if book.Title != "Hard Book" {
		t.Errorf("book.Title = %s; excepted 'Hard Book' ", book.Title)
	}

	book, _ = lib.Get("Special Book")
	if book.Title != "Special Book" {
		t.Errorf("book.Title = %s; excepted 'Special Book' ", book.Title)
	}
}

func TestSliceLibrary(t *testing.T) {
	sliceStorage := storage.NewSliceStorage()
	lib := library.NewLibrary(sliceStorage)

	boringBook := Book{Title: "Boring Book"}
	stupidBook := Book{Title: "Stupid Book"}
	fastBook := Book{Title: "Fast Book"}
	shortBook := Book{Title: "Short Book"}
	beautifulBook := Book{Title: "Beautiful Book"}
	crashedBook := Book{Title: "Crashed Book"}
	longBook := Book{Title: "Long Book"}
	loveBook := Book{Title: "Love Book", Content: "<3"}

	lib.Add(boringBook)
	lib.Add(stupidBook)
	lib.Add(fastBook)
	lib.Add(shortBook)
	lib.Add(beautifulBook)
	lib.Add(loveBook)

	if lib.Size() != 6 {
		t.Errorf("Lib.Size() = %d; excepted 6", lib.Size())
	}

	book, _ := lib.Get("Boring Book")
	if book.Title != "Boring Book" {
		t.Errorf("book.Title = %s; excepted 'Boring Book' ", book.Title)
	}

	book, _ = lib.Get("Love Book")
	if book.Content != "<3" {
		t.Errorf("book.Content = %s; excepted '<3' ", book.Content)
	}

	book, _ = lib.Get("Short Book")
	id := book.ID()
	lib.Remove(id)

	if lib.Size() != 5 {
		t.Errorf("Lib.Size() = %d; excepted 5", lib.Size())
	}

	book, ok := lib.Get("Short Book")
	if ok {
		t.Errorf("book.Title = %s, ok = true; excepted '', ok = false ", book.Title)
	}

	lib.SetIdGenerator(identification.AlmostDefaultGenerator())
	lib.Add(crashedBook)
	lib.Add(longBook)

	book, _ = lib.Get("Crashed Book")
	if book.Title != "Crashed Book" {
		t.Errorf("book.Title = %s; excepted 'Crashed Book' ", book.Title)
	}

	book, _ = lib.Get("Long Book")
	if book.Title != "Long Book" {
		t.Errorf("book.Title = %s; excepted 'Long Book' ", book.Title)
	}
}
