package models

// Country structure on database
type Country struct {
	DefaultData
}

func CreateCountry() Country {
	return Country{}
}

func ExtractCountry(c Country) []any {
	return []any{c.Id, c.Name}
}

func RecoveryCountry(c *Country, data ...any) {
	RecoveryData(&c.DefaultData, data...)
}
