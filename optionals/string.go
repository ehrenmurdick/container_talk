package optionals

import (
	"errors"
	"fmt"
	"math/rand"
)

type OptionalString interface {
	FlatMap(func(string) (string, error)) OptionalString
	HandleErr(func(error) error) OptionalString
	Print() OptionalString
}

type someString struct {
	value string
}

type noneString struct {
	err error
}

func String(s string, e error) OptionalString {
	if e != nil {
		return noneString{
			err: e,
		}
	}

	return someString{
		value: s,
	}
}

func (s someString) FlatMap(f func(string) (string, error)) OptionalString {
	return String(f(s.value))
}

func (n noneString) FlatMap(f func(string) (string, error)) OptionalString {
	return n
}

func (s someString) HandleErr(f func(error) error) OptionalString {
	return s
}

func (n noneString) HandleErr(f func(error) error) OptionalString {
	return noneString{
		err: f(n.err),
	}
}

func (s someString) Print() OptionalString {
	return s.FlatMap(func(str string) (string, error) {
		err := maybeError()
		if err == nil {
			println(str)
			return str, nil
		} else {
			return "", err
		}
	})
}

func (n noneString) Print() OptionalString {
	return n
}

var tries int = 0

func maybeError() error {
	tries++
	if rand.Intn(10) > 5 {
		return errors.New(fmt.Sprintf("failed on attempt %v", tries))
	} else {
		return nil
	}
}
