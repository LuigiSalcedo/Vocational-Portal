package models

// Default data struct
type DefaultData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func RecoveryData(d *DefaultData, data ...any) {
	d.Id = data[0].(int64)
	d.Name = data[1].(string)
}
