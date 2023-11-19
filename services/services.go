package services

import (
	"slices"
	"strings"
	"vocaportal/models"
)

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

// Sort a slice using the position of the prefix given
func SortWithPrefix[T models.Sortable](slice []T, prefix string) {
	slices.SortFunc(slice, func(v1, v2 T) int {
		p1 := strings.Index(v1.ByName(), prefix)
		p2 := strings.Index(v2.ByName(), prefix)

		if p1 != p2 {
			return p1 - p2
		}

		return strings.Compare(v1.ByName(), v2.ByName())
	})
}

// This function create a new SimpleFilter and add it to a filter slice
func CreateAndAddFilter[T any](filters []Filter[T], ef func(v *T) any, cmpValue int64) []Filter[T] {
	if cmpValue != -1 {
		filters = append(filters, NewSimpleFilter(func(v *T) any {
			return ef(v)
		}, cmpValue))
	}

	return filters
}
