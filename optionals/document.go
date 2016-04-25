package optionals

import (
	"github.com/ehrenmurdick/container_talk/entities"
)

type OptionalDocument interface {
	Try(func(entities.Document) (entities.Document, error)) OptionalDocument
	HandleErr(func(error) error) OptionalDocument
	PrintErr() OptionalDocument
	ToString() OptionalString

	Print() OptionalDocument
	Save() OptionalDocument
}

type SomeDocument struct {
	value entities.Document
}

type NoneDocument struct {
	err error
}

func WrapDocument(s entities.Document, e error) OptionalDocument {
	if e != nil {
		return NoneDocument{
			err: e,
		}
	}

	return SomeDocument{
		value: s,
	}
}

func (s SomeDocument) Try(f func(entities.Document) (entities.Document, error)) OptionalDocument {
	return WrapDocument(f(s.value))
}

func (n NoneDocument) Try(f func(entities.Document) (entities.Document, error)) OptionalDocument {
	return n
}

func (s SomeDocument) HandleErr(f func(error) error) OptionalDocument {
	return s
}

func (n NoneDocument) HandleErr(f func(error) error) OptionalDocument {
	return NoneDocument{
		err: f(n.err),
	}
}

func (s SomeDocument) PrintErr() OptionalDocument {
	return s
}

func (n NoneDocument) PrintErr() OptionalDocument {
	return n.HandleErr(func(err error) error {
		println(err.Error())
		return err
	})
}

func (s SomeDocument) ToString() OptionalString {

	return WrapString(s.value.ToString(), nil)

}

func (n NoneDocument) ToString() OptionalString {
	return NoneString{
		err: n.err,
	}
}
func (s SomeDocument) Print() OptionalDocument {
	return s.Try(func(d entities.Document) (entities.Document, error) {
		return d, d.Print()
	})
}

func (n NoneDocument) Print() OptionalDocument {
	return n
}

func (s SomeDocument) Save() OptionalDocument {
	return s.Try(func(d entities.Document) (entities.Document, error) {
		return d, d.Save()
	})
}

func (n NoneDocument) Save() OptionalDocument {
	return n
}
