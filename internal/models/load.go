package models

// Load represents a single load, holding the pickup and dropoff locations
// and keeping track of whether or not it has been completed
type Load struct {
	number   int
	Pickup   *Location
	Dropoff  *Location
	complete bool
}

// NewLoad is a factory function used in creating a LoadSet when reading a problem file
func NewLoad(num int, pickup *Location, dropoff *Location, complete bool) *Load {
	return &Load{
		number:   num,
		Pickup:   pickup,
		Dropoff:  dropoff,
		complete: complete,
	}
}

// clone just makes a Load that has not been completed, regardless
// of the completion status of the original
func (l *Load) clone() *Load {
	return &Load{
		number:  l.number,
		Pickup:  l.Pickup,
		Dropoff: l.Dropoff,
	}
}
