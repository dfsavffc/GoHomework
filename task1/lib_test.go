package main

import (
	"GoHomework/task1/book"
	"GoHomework/task1/identification"
	"GoHomework/task1/library"
	"GoHomework/task1/storage"
	"testing"
)

func TestMapLibrary(t *testing.T) {
	mapStorage := storage.NewMapStorage()
	lib := library.NewLibrary(mapStorage)

	justBook := *book.NewBook("Just Book", "", "")
	simpleBook := *book.NewBook("Simple Book", "", "")
	normalBook := *book.NewBook("Normal Book", "", "")
	hardBook := *book.NewBook("Hard Book", "", "")
	specialBook := *book.NewBook("Special Book", "", "")
	expensiveBook := *book.NewBook("Expensive Book", "?", "")
	goodBook := *book.NewBook("Good Book", "", "")
	wonderfulBook := *book.NewBook("Wonderful Book", "", "")

	lib.Add(justBook)
	lib.Add(simpleBook)
	lib.Add(normalBook)
	lib.Add(wonderfulBook)
	lib.Add(expensiveBook)
	lib.Add(goodBook)

	if lib.Size() != 6 {
		t.Errorf("Lib.Size() = %d; excepted 6", lib.Size())
	}

	eBook, ok := lib.Get("Simple Book")
	if ok != nil {
		t.Errorf("lib.Get('Simple Book') = %v; excepted nil", ok)
	}
	if eBook.Title != "Simple Book" {
		t.Errorf("eBook.Title = %s; excepted 'Simple Book' ", eBook.Title)
	}

	eBook, ok = lib.Get("Expensive Book")
	if ok != nil {
		t.Errorf("lib.Get('Expensive Book') = %v; excepted nil", ok)
	}
	if eBook.Content != "?" {
		t.Errorf("eBook.Content = %s; excepted '?' ", eBook.Content)
	}

	eBook, _ = lib.Get("Expensive Book")
	id := eBook.ID()
	ok = lib.Remove(id)
	if ok != nil {
		t.Errorf("lib.Remove(id) = %v; excepted nil", ok)
	}

	if lib.Size() != 5 {
		t.Errorf("Lib.Size() = %d; excepted 5", lib.Size())
	}

	eBook, ok = lib.Get("Expensive Book")
	if ok == nil {
		t.Errorf("eBook.Title = %s, ok = nil; excepted '', ok = error ", eBook.Title)
	}

	lib.SetIdGenerator(identification.AlmostDefaultGenerator())
	lib.Add(hardBook)
	lib.Add(specialBook)

	eBook, ok = lib.Get("Hard Book")
	if ok != nil {
		t.Errorf("lib.Get('Hard Book') = %v; excepted nil", ok)
	}
	if eBook.Title != "Hard Book" {
		t.Errorf("eBook.Title = %s; excepted 'Hard Book' ", eBook.Title)
	}

	eBook, ok = lib.Get("Special Book")
	if ok != nil {
		t.Errorf("lib.Get('Special Book') = %v; excepted nil", ok)
	}
	if eBook.Title != "Special Book" {
		t.Errorf("eBook.Title = %s; excepted 'Special Book' ", eBook.Title)
	}
}

func TestSliceLibrary(t *testing.T) {
	sliceStorage := storage.NewSliceStorage()
	lib := library.NewLibrary(sliceStorage)

	boringBook := *book.NewBook("Boring Book", "", "")
	stupidBook := *book.NewBook("Stupid Book", "", "")
	fastBook := *book.NewBook("Fast Book", "", "")
	shortBook := *book.NewBook("Short Book", "", "")
	beautifulBook := *book.NewBook("Beautiful Book", "", "")
	crashedBook := *book.NewBook("Crashed Book", "", "")
	longBook := *book.NewBook("Long Book", "", "")
	loveBook := *book.NewBook("Love Book", "<3", "")

	lib.Add(boringBook)
	lib.Add(stupidBook)
	lib.Add(fastBook)
	lib.Add(shortBook)
	lib.Add(beautifulBook)
	lib.Add(loveBook)

	if lib.Size() != 6 {
		t.Errorf("Lib.Size() = %d; excepted 6", lib.Size())
	}

	eBook, ok := lib.Get("Boring Book")
	if ok != nil {
		t.Errorf("lib.Get('Boring Book') = %v; excepted nil", ok)
	}
	if eBook.Title != "Boring Book" {
		t.Errorf("eBook.Title = %s; excepted 'Boring Book' ", eBook.Title)
	}

	eBook, ok = lib.Get("Love Book")
	if ok != nil {
		t.Errorf("lib.Get('Love Book') = %v; excepted nil", ok)
	}
	if eBook.Content != "<3" {
		t.Errorf("eBook.Content = %s; excepted '<3' ", eBook.Content)
	}

	eBook, _ = lib.Get("Short Book")
	id := eBook.ID()
	ok = lib.Remove(id)
	if ok != nil {
		t.Errorf("lib.Remove(id) = %v; excepted nil", ok)
	}

	if lib.Size() != 5 {
		t.Errorf("Lib.Size() = %d; excepted 5", lib.Size())
	}

	eBook, ok = lib.Get("Short Book")
	if ok == nil {
		t.Errorf("eBook.Title = %s, ok = true; excepted '', ok = false ", eBook.Title)
	}

	lib.SetIdGenerator(identification.AlmostDefaultGenerator())
	lib.Add(crashedBook)
	lib.Add(longBook)

	eBook, ok = lib.Get("Crashed Book")
	if ok != nil {
		t.Errorf("lib.Get('Crashed Book') = %v; excepted nil", ok)
	}
	if eBook.Title != "Crashed Book" {
		t.Errorf("eBook.Title = %s; excepted 'Crashed Book' ", eBook.Title)
	}

	eBook, ok = lib.Get("Long Book")
	if ok != nil {
		t.Errorf("lib.Get('Long Book') = %v; excepted nil", ok)
	}
	if eBook.Title != "Long Book" {
		t.Errorf("eBook.Title = %s; excepted 'Long Book' ", eBook.Title)
	}
}
