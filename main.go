package main

import "errors"

//go:generate ./optional Document entities.Document github.com/ehrenmurdick/container_talk/entities Print Save
//go:generate ./optional String string

import (
	"github.com/ehrenmurdick/container_talk/entities"
	"github.com/ehrenmurdick/container_talk/optionals"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	doc0 := entities.NewDocument("document 1")
	doc0, err := doc0.SetContent("not long")
	if err != nil {
		printError(err)
	} else {
		_, err = printDocument(doc0)
		if err != nil {
			printError(err)
		} else {
			_, err = saveDocument(doc0)
			if err != nil {
				printError(err)
			}
		}
	}

	doc1 := entities.NewDocument("document 2")
	opt1 := optionals.WrapAny(doc1, nil)
	opt1.
		Try(printAny).
		Try(saveAny).
		HandleErr(printError)

	doc2 := entities.NewDocument("document 3")
	opt2 := optionals.WrapDocument(doc2, nil)
	opt2.
		Try(printDocument).
		Try(saveDocument).
		HandleErr(printError)
}

func printDocument(d entities.Document) (entities.Document, error) {
	println(d.Content())
	return d, nil
}

func saveDocument(d entities.Document) (entities.Document, error) {
	err := d.Save()
	return d, err
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
