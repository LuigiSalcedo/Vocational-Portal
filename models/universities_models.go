package models

// University data on database
type University struct {
	DefaultData
	City City `json:"ciudad"`
}

func CreateUniversity() *University {
	return &University{}
}

func (u *University) Extract() []any {
	data := make([]any, 0, 6)
	data = append(data, u.Id, u.Name)
	data = append(data, u.City.Extract()...)
	return data
}

func (u *University) Recovery(data ...any) {
	u.DefaultData.Recovery(data...)
	u.City.Recovery(data[2:]...)
}
