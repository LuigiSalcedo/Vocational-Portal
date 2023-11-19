package models

// Academic offer database model
type Offer struct {
	DefaultData
	Programm   Programm   `json:"programa"`
	University University `json:"universidad"`
}

// Create an offer
func CreateOffer() *Offer {
	return &Offer{}
}

// Extract func
func (o *Offer) Extract() []any {
	data := make([]any, 0, 8)

	data = append(data, o.DefaultData.Extract()...)
	data = append(data, o.Programm.Extract()...)
	data = append(data, o.University.Extract()...)

	return data
}

// Recovery func
func (o *Offer) Recovery(data ...any) {
	o.DefaultData.Recovery(data[:2]...)
	o.Programm.Recovery(data[2:4]...)
	o.University.Recovery(data[4:]...)
}
