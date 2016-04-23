package main

//go:generate ./optional Document entities.Document github.com/ehrenmurdick/container_talk/entities Print Save

//go:generate ./optional String string

import (
	"fmt"
	"github.com/ehrenmurdick/container_talk/entities"
	"github.com/ehrenmurdick/container_talk/optionals"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	doc := entities.NewDocument("hello, world!")
	opt := optionals.WrapDocument(doc, nil)

	opt.
		Print().
		Save().
		FlatMap(printLen).
		PrintErr()
}

func printLen(d entities.Document) (entities.Document, error) {
	fmt.Printf("document is %v bytes in length\n", len(d.Content()))
	return d, nil
}
