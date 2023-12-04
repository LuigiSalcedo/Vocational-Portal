package models

// Academic offer database model
type Offer struct {
	DefaultData
	Price      string     `json:"precio"`
	URL        string     `json:"url"`
	Programm   Programm   `json:"programa"`
	University University `json:"universidad"`
}

// Create an offer
func CreateOffer() *Offer {
	return &Offer{}
}

// Extract func
func (o *Offer) Extract() []any {
	data := make([]any, 0, 10)

	data = append(data, o.DefaultData.Extract()...)
	data = append(data, o.Price, o.URL)
	data = append(data, o.Programm.Extract()...)
	data = append(data, o.University.Extract()...)

	return data
}

// Recovery func
func (o *Offer) Recovery(data ...any) {
	o.DefaultData.Recovery(data[:2]...)
	o.Price = data[2].(string)
	o.URL = data[3].(string)
	o.Programm.Recovery(data[4:7]...)
	o.University.Recovery(data[7:]...)
}
