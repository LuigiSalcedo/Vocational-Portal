package models

// Study area database model
type Area struct {
	DefaultData
}

func CreateArea() *Area {
	return &Area{}
}
