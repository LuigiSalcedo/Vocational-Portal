package models

// City database struct
type City struct {
	DefaultData
	Country Country `json:"pais"`
}

// Creator func
func CreateCity() *City {
	return &City{}
}

// Extractor Func
func (c *City) Extract() []any {
	data := make([]any, 0, 4)
	data = append(data, c.Id, c.Name)
	data = append(data, c.Country.Extract()...)
	return data
}

// Recovery
func (c *City) Recovery(data ...any) {
	c.Id = data[0].(int64)
	c.Name = data[1].(string)
	c.Country.Recovery(data[2:]...)
}
