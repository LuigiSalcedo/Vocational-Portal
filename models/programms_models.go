package models

import (
	"strconv"
)

// Academic programm database model
type Programm struct {
	DefaultData
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

// Create a program with area relation
func CreatePWAR() *PWAR {
	return &PWAR{}
}

func (p *PWAR) Extract() []any {
	data := make([]any, 0, 3)
	data = append(data, p.Programm.Extract()...)
	data = append(data, p.N)
	return data
}

func (p *PWAR) Recovery(data ...any) {
	p.Programm.Recovery(data[:2]...)
	n, err := strconv.ParseFloat(string(data[2].([]uint8)), 64)

	if err != nil {
		panic(err)
	}

	p.N = n
}
