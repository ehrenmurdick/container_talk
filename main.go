package main

import (
	"./entities"
	"./optionals"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	doc := entities.NewDocument("hello, world!")
	opt := optionals.WrapDocument(doc, nil)

	for x := 0; x < 10; x++ {
		opt = opt.Print()
	}

	opt.HandleErr(func(e error) error {
		println(e.Error())
		return e
	})
}
