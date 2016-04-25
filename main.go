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

	doc1 := entities.NewDocument("hello, world!")
	opt1 := optionals.WrapAny(doc1, nil)
	op1.
		FlatMap(printAny).
		FlatMap(saveAny).
		HandleErr(HandleErr)

	doc2 := entities.NewDocument("hello, world!")
	opt2 := optionals.WrapDocument(doc2, nil)

	opt2.
		FlatMap(printDocument).
		FlatMap(saveDocument).
		HandleErr(handleError)
}

func printDocument(d entities.Document) (entities.Document, error) {
	println(d.Content())
	return d, nil
}

func saveDocument(d entities.Document) (entities.Document, error) {
	err := d.Save()
	return d, err
}

func handleError(err error) error {
	println(err.Error())
	return err
}

func saveAny(i interface{}) (interface{}, error) {
	var doc entities.Document
	doc, ok := i.(entities.Document)

	if ok {
		doc.Save()
		return doc, nil
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
