package main

import (
	"errors"
	"github.com/ehrenmurdick/container_talk/optionals"
)

func p(s *string) *string {
	print(*s)
	return s
}

var willFail bool = false

func couldFail() (string, error) {
	if willFail {
		return "", errors.New("failed")
	} else {
		willFail = true
		return "success", nil
	}
}

func main() {
	var c optionals.OptionalString

	c = optionals.String(nil)
	c.FlatMap(p)

	var s string = "hello\n"
	c = optionals.String(&s)
	c.FlatMap(p)
}
