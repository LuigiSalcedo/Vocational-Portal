package models

// Country structure on database
type Country struct {
	DefaultData
}

func CreateCountry() *Country {
	return &Country{}
}

func (c *Country) Extract() []any {
	return []any{c.Id, c.Name}
}

func (c *Country) Recovery(data ...any) {
	c.DefaultData.Recovery(data...)
}
