package services

// Filter is a func that the register should pass to be returned usually
// the filter is a comparation with a query param
type Filter[T any] func(v *T) bool

// Function to filter data
func FilterData[T any](data []*T, filters ...Filter[T]) []*T {
	result := make([]*T, 0, len(data))

	for _, v := range data {
		if passFilters(v, filters...) {
			result = append(result, v)
		}
	}

	return result
}

// Function to create a filter
func NewSimpleFilter[T any](ef func(strct *T) any, value any) Filter[T] {
	return func(u *T) bool {
		return ef(u) == value
	}
}

func passFilters[T any](v *T, filters ...Filter[T]) bool {
	for _, filter := range filters {
		if !filter(v) {
			return false
		}
	}
	return true
}
