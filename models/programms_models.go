package models

import (
	"strconv"
)

// Academic programm database model
type Programm struct {
	DefaultData
	Description string `json:"descripcion"`
}

// PWAR = Program With Area Relation
type PWAR struct {
	Programm
	N float64
}

// Create a Program
func CreateProgramm() *Programm {
	return &Programm{}
}

func (p *Programm) Extract() []any {
	data := make([]any, 0, 3)
	data = append(data, p.DefaultData.Extract()...)
	data = append(data, p.Description)
	return data
}

func (p *Programm) Recovery(data ...any) {
	p.DefaultData.Recovery(data[:2]...)
	p.Description = data[2].(string)
}

// Create a program with area relation
func CreatePWAR() *PWAR {
	return &PWAR{}
}

func (p *PWAR) Extract() []any {
	data := make([]any, 0, 4)
	data = append(data, p.Programm.Extract()...)
	data = append(data, p.N)
	return data
}

func (p *PWAR) Recovery(data ...any) {
	p.Programm.Recovery(data[:3]...)
	n, err := strconv.ParseFloat(string(data[3].([]uint8)), 64)

	if err != nil {
		panic(err)
	}

	p.N = n
}
