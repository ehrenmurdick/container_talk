package main

import (
	"./optionals"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	c := optionals.String("hello, world!", nil)
	for x := 0; x < 10; x++ {
		c = c.Print()
	}

	c.HandleErr(func(e error) error {
		println(e.Error())
		return e
	})
}
