package models

// City database struct
type City struct {
	DefaultData
	Country Country `json:"pais"`
}

// Creator func
func CreateCity() City {
	return City{}
}

// Extractor Func
func ExtractCity(c City) []any {
	data := make([]any, 0, 4)
	data = append(data, c.Id, c.Name)
	data = append(data, ExtractCountry(c.Country)...)
	return data
}

// Recovery
func RecoveryCity(c *City, data ...any) {
	c.Id = data[0].(int64)
	c.Name = data[1].(string)
	RecoveryCountry(&c.Country, data[2:]...)
}
