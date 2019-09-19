package compare

type IntComparator func(int) bool

func LessThanInt(value int) IntComparator {
	return func(compare int) bool {
		return compare < value
	}
}

func LessThanOrEqualInt(value int) IntComparator {
	return func(compare int) bool {
		return compare <= value
	}
}

func EqualInt(value int) IntComparator {
	return func(compare int) bool {
		return compare == value
	}
}

func NotEqualInt(value int) IntComparator {
	return func(compare int) bool {
		return compare != value
	}
}

func GreaterThanOrEqualInt(value int) IntComparator {
	return func(compare int) bool {
		return compare >= value
	}
}

func GreaterThanInt(value int) IntComparator {
	return func(compare int) bool {
		return compare > value
	}
}
