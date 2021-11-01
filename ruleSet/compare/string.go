package compare

type StringComparator func(string) bool

func EqualString(value string) StringComparator {
	return func(compare string) bool {
		if compare == "" {
			return false
		}
		return compare == value
	}
}

func NotEqualString(value string) StringComparator {
	return func(compare string) bool {
		if compare == "" {
			return false
		}
		return compare != value
	}
}
