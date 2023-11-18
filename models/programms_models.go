package models

// Academic programm database model
type Programm struct {
	DefaultData
}

// Create a Program
func CreateProgramm() *Programm {
	return &Programm{}
}

// Extract func
func (p *Programm) Extract() []any {
	return p.DefaultData.Extract()
}

// Recovery func
func (p *Programm) Recovery(data ...any) {
	p.DefaultData.Recovery(data...)
}
