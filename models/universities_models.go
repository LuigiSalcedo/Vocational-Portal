package models

// University data on database
type University struct {
	DefaultData
}

func CreateUniversity() University {
	return University{}
}

func ExtractUniversity(u University) []any {
	return []any{u.Id, u.Name}
}

func RecoveryUniversity(u *University, data ...any) {
	RecoveryData(&u.DefaultData, data...)
}
