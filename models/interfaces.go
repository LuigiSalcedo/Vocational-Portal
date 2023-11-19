package models

// Interface that should implement everymodel that will be used as database query result.
type DatabaseEntity interface {
	Recovery(data ...any)
	Extract() []any
}

// Interface used to sort in a slices
type Sortable interface {
	ByName() string
}
