package models

// Default data struct
type DefaultData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Interface that should implement everymodel that will be used as database query result.
type DatabaseEntity interface {
	Recovery(data ...any)
	Extract() []any
}

func (d *DefaultData) Recovery(data ...any) {
	d.Id = data[0].(int64)
	d.Name = data[1].(string)
}

func (d *DefaultData) Extract() []any {
	return []any{d.Id, d.Name}
}
