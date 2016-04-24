package optionals

type OptionalString interface {
	FlatMap(func(string) (string, error)) OptionalString
	HandleErr(func(error) error) OptionalString
	PrintErr() OptionalString
	ToString() OptionalString
}

type SomeString struct {
	value string
}

type NoneString struct {
	err error
}

func WrapString(s string, e error) OptionalString {
	if e != nil {
		return NoneString{
			err: e,
		}
	}

	return SomeString{
		value: s,
	}
}

func (s SomeString) FlatMap(f func(string) (string, error)) OptionalString {
	return WrapString(f(s.value))
}

func (n NoneString) FlatMap(f func(string) (string, error)) OptionalString {
	return n
}

func (s SomeString) HandleErr(f func(error) error) OptionalString {
	return s
}

func (n NoneString) HandleErr(f func(error) error) OptionalString {
	return NoneString{
		err: f(n.err),
	}
}

func (s SomeString) PrintErr() OptionalString {
	return s
}

func (n NoneString) PrintErr() OptionalString {
	return n.HandleErr(func(err error) error {
		println(err.Error())
		return err
	})
}

func (s SomeString) ToString() OptionalString {

	return WrapString(s.value, nil)

}

func (n NoneString) ToString() OptionalString {
	return NoneString{
		err: n.err,
	}
}
