package models

// University data on database
type University struct {
	DefaultData
	City City   `json:"ciudad"`
	URL  string `json:"url"`
}

func CreateUniversity() *University {
	return &University{}
}

func (u *University) Extract() []any {
	data := make([]any, 0, 6)
	data = append(data, u.Id, u.Name, u.URL)
	data = append(data, u.City.Extract()...)
	return data
}

func (u *University) Recovery(data ...any) {
	u.DefaultData.Recovery(data[:2]...)
	u.URL = data[2].(string)
	u.City.Recovery(data[3:]...)
}
