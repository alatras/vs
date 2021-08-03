package compare

type Float64Comparator func(float64) bool

func LessThanFloat64(value float64) Float64Comparator {
	return func(compare float64) bool {
		return compare < value
	}
}

func LessThanOrEqualFloat64(value float64) Float64Comparator {
	return func(compare float64) bool {
		return compare <= value
	}
}

func EqualFloat64(value float64) Float64Comparator {
	return func(compare float64) bool {
		return compare == value
	}
}

func NotEqualFloat64(value float64) Float64Comparator {
	return func(compare float64) bool {
		return compare != value
	}
}

func GreaterThanOrEqualFloat64(value float64) Float64Comparator {
	return func(compare float64) bool {
		return compare >= value
	}
}

func GreaterThanFloat64(value float64) Float64Comparator {
	return func(compare float64) bool {
		return compare > value
	}
}
