package optionals

type OptionalString interface {
	FlatMap(func(string) (string, error)) OptionalString
	HandleErr(func(error) error) OptionalString
	PrintErr() OptionalString
}

type someString struct {
	value string
}

type noneString struct {
	err error
}

func WrapString(s string, e error) OptionalString {
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
	return WrapString(f(s.value))
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

func (s someString) PrintErr() OptionalString {
	return s
}

func (n noneString) PrintErr() OptionalString {
	return n.HandleErr(func(err error) error {
		println(err.Error())
		return err
	})
}
