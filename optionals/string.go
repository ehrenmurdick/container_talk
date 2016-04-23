package optionals

type OptionalString interface {
	FlatMap(func(*string) *string) OptionalString
}

type someString struct {
	value *string
}

type noneString struct{}

func String(s *string) OptionalString {
	if s == nil {
		return noneString{}
	}

	return someString{
		value: s,
	}
}

func (s someString) FlatMap(f func(*string) *string) OptionalString {
	return String(f(s.value))
}

func (s noneString) FlatMap(f func(*string) *string) OptionalString {
	return noneString{}
}
