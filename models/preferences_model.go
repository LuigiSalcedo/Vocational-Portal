package models

// Preference database entity
type Preference struct {
	DefaultData
}

// Create a preferences register
func CreatePreference() (p *Preference) {
	return &Preference{}
}
