package optionals



type OptionalString interface {
	Try(func(string) (string, error)) OptionalString
	HandleErr(func(error) error) OptionalString

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

func (s SomeString) Try(f func(string) (string, error)) OptionalString {
	return WrapString(f(s.value))
}

func (n NoneString) Try(f func(string) (string, error)) OptionalString {
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

