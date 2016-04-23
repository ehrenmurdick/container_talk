package main

import "github.com/ehrenmurdick/container_talk/optionals"

func main() {
	c := optionals.String("Hello world!\n")

	c.FlatMap(func(s string) string {
		print(s)
		return s
	})
}
