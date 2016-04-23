package main

//go:generate go generate entities/document.go

import (
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
		PrintErr()
}
