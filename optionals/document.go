package optionals

import (
	"github.com/ehrenmurdick/container_talk/entities"
)

type OptionalDocument interface {
	FlatMap(func(entities.Document) (entities.Document, error)) OptionalDocument
	HandleErr(func(error) error) OptionalDocument
	Print() OptionalDocument
	Save(filename string) OptionalDocument
	PrintErr() OptionalDocument
}

type someDocument struct {
	value entities.Document
}

type noneDocument struct {
	err error
}

func WrapDocument(s entities.Document, e error) OptionalDocument {
	if e != nil {
		return noneDocument{
			err: e,
		}
	}

	return someDocument{
		value: s,
	}
}

func (s someDocument) FlatMap(f func(entities.Document) (entities.Document, error)) OptionalDocument {
	return WrapDocument(f(s.value))
}

func (n noneDocument) FlatMap(f func(entities.Document) (entities.Document, error)) OptionalDocument {
	return n
}

func (s someDocument) HandleErr(f func(error) error) OptionalDocument {
	return s
}

func (n noneDocument) HandleErr(f func(error) error) OptionalDocument {
	return noneDocument{
		err: f(n.err),
	}
}

func (s someDocument) PrintErr() OptionalDocument {
	return s
}

func (n noneDocument) PrintErr() OptionalDocument {
	return n.HandleErr(func(err error) error {
		println(err.Error())
		return err
	})
}

func (s someDocument) Print() OptionalDocument {
	return s.FlatMap(func(d entities.Document) (entities.Document, error) {
		return d, d.Print()
	})
}

func (n noneDocument) Print() OptionalDocument {
	return n
}

func (s someDocument) Save(filename string) OptionalDocument {
	return s.FlatMap(func(d entities.Document) (entities.Document, error) {
		return d, d.Save(filename)
	})
}

func (n noneDocument) Save(filename string) OptionalDocument {
	return n
}
