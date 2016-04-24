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
		ToString().
		FlatMap(printLen).
		FlatMap(echo).
		PrintErr()
}

func echo(str string) (string, error) {
	println(str)
	return str, nil
}

func printLen(str string) (string, error) {
	fmt.Printf("string is %v bytes in length\n", len(str))
	return str, nil
}
