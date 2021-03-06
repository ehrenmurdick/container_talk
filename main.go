package main

import "errors"

//go:generate ./optional Document Print Save

import (
	"github.com/ehrenmurdick/container_talk/entities"
	"github.com/ehrenmurdick/container_talk/optionals"
)

func main() {
	doc0 := entities.NewDocument("document 1")
	_, err := printDocument(doc0)
	if err != nil {
		printError(err)
	} else {
		_, err = saveDocument(doc0)
		if err != nil {
			printError(err)
		}
	}

	println("\n\n")

	doc1 := entities.NewDocument("document 2")
	opt1 := optionals.WrapAny(doc1, nil)
	opt1.
		Try(printAny).
		Try(saveAny).
		HandleErr(printError)

	println("\n\n")

	doc2 := entities.NewDocument("document 3")
	opt2 := optionals.WrapDocument(doc2, nil)
	opt2.
		Try(printDocument).
		Try(saveDocument).
		HandleErr(printError)

	println("\n\n")

	doc3 := entities.NewDocument("document 4")
	opt3 := optionals.WrapDocument(doc3, nil)
	opt3.
		Print().
		Save().
		HandleErr(printError)
}

func renameAny(i interface{}) (interface{}, error) {
	doc, ok := i.(entities.Document)
	if !ok {
		return nil, errors.New("not a document")
	}

	return doc.SetContent("too long of a document content")
}

func printDocument(d entities.Document) (entities.Document, error) {
	println(d.Content())
	return d, nil
}

func saveDocument(d entities.Document) (entities.Document, error) {
	return d, d.Save()
}

func printError(err error) error {
	println(err.Error())
	return err
}

func saveAny(i interface{}) (interface{}, error) {
	var doc entities.Document
	doc, ok := i.(entities.Document)

	if ok {
		return doc, doc.Save()
	} else {
		return nil, errors.New("not a document")
	}
}

func printAny(i interface{}) (interface{}, error) {
	var doc entities.Document
	doc, ok := i.(entities.Document)

	if ok {
		doc.Print()
		return doc, nil
	} else {
		return nil, errors.New("not a document")
	}
}

func renameDocument(d entities.Document) (entities.Document, error) {
	return d.SetContent("too long of a string")
}
