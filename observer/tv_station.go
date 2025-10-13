package observer

import "fmt"

// TvStation represents the TV channel broadcaster, that could be subscribed to the specific TV shows.
type TvStation struct {
	name      string
	frequency float32
}

// Update implements the callback function for the external publishers (TV shows).
func (t *TvStation) Update(data any) {
	fmt.Printf("[%.2f MHz | %v]: %v\n", t.frequency, t.name, data)
}

// NewTvStation creates the new TV station.
func NewTvStation(name string, frequency float32) *TvStation {
	return &TvStation{name, frequency}
}

// Interface implementation assertion.
var _ Observer = &TvStation{}
