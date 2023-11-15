package models

// Univerisity data on database
type Univerisity struct {
	DefaultData
}

func CreateUniversity() Univerisity {
	return Univerisity{}
}

func ExtractUniversity(u Univerisity) []any {
	return []any{u.Id, u.Name}
}

func RecoveryUniversity(u *Univerisity, args ...any) {
	u.Id = args[0].(int64)
	u.Name = args[1].(string)
}
