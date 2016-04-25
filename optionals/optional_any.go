package optionals

type OptionalAny interface {
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
