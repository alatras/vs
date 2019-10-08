package compare

type StringComparator func(string) bool

func EqualString(value string) StringComparator {
	return func(compare string) bool {
		return compare == value
	}
}

func NotEqualString(value string) StringComparator {
	return func(compare string) bool {
		return compare != value
	}
}
