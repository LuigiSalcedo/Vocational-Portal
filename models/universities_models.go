package models

// University data on database
type University struct {
	DefaultData
}

func CreateUniversity() *University {
	return &University{}
}

func (u *University) Extract() []any {
	return []any{u.Id, u.Name}
}

func (u *University) Recovery(data ...any) {
	u.DefaultData.Recovery(data...)
}
