package optionals

type OptionalAny interface {
	FlatMap(f func(interface{}) (interface{}, error)) OptionalAny
	HandleErr(f func(error) error) OptionalAny
}

type SomeAny struct {
	value interface{}
}

type NoneAny struct {
	err error
}

func WrapAny(s interface{}, e error) OptionalAny {
	if e != nil {
		return NoneAny{
			err: e,
		}
	}
	return SomeAny{
		value: s,
	}
}

func (s SomeAny) FlatMap(f func(interface{}) (interface{}, error)) OptionalAny {
	return WrapAny(f(s.value))
}

func (n NoneAny) FlatMap(f func(interface{}) (interface{}, error)) OptionalAny {
	return n
}

func (s SomeAny) HandleErr(f func(error) error) OptionalAny {
	return s
}

func (n NoneAny) HandleErr(f func(error) error) OptionalAny {
	return NoneAny{
		err: f(n.err),
	}
}
