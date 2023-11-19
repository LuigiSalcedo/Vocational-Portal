package models

// Default data struct
type DefaultData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (d *DefaultData) Recovery(data ...any) {
	d.Id = data[0].(int64)
	d.Name = data[1].(string)
}

func (d *DefaultData) Extract() []any {
	return []any{d.Id, d.Name}
}

func (d *DefaultData) ByName() string {
	return d.Name
}
