package compare

type Uint64Comparator func(uint64) bool

func LessThanUint64(value uint64) Uint64Comparator {
	return func(compare uint64) bool {
		return compare < value
	}
}

func LessThanOrEqualUint64(value uint64) Uint64Comparator {
	return func(compare uint64) bool {
		return compare <= value
	}
}

func EqualUint64(value uint64) Uint64Comparator {
	return func(compare uint64) bool {
		return compare == value
	}
}

func NotEqualUint64(value uint64) Uint64Comparator {
	return func(compare uint64) bool {
		return compare != value
	}
}

func GreaterThanOrEqualUint64(value uint64) Uint64Comparator {
	return func(compare uint64) bool {
		return compare >= value
	}
}

func GreaterThanUint64(value uint64) Uint64Comparator {
	return func(compare uint64) bool {
		return compare > value
	}
}
