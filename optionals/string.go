package optionals

type OptionalString interface {
	FlatMap(func(string) string) OptionalString
}

type optionalString struct {
	value string
}

func String(s string) OptionalString {
	return optionalString{
		value: s,
	}
}

func (s optionalString) FlatMap(f func(string) string) OptionalString {
	return String(f(s.value))
}
