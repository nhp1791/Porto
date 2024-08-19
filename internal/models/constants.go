// Package models defines the structs that hold information about the scheduling problem
// as well as those that are used to help solve it
package models

const (
	lparen                     = '('
	rparen                     = ')'
	maxDriverHours      uint64 = 12
	maxDriverMinutes    uint64 = 60
	maxDriverSqDistance uint64 = maxDriverHours * maxDriverHours * maxDriverMinutes * maxDriverMinutes
	// MaxNearestNeighbors controls how many possible paths are tested in the solution.  If this
	// value is larger than the size of the load set, then the size of the load set will be used.
	MaxNearestNeighbors        = 10
	costPerDriver       uint64 = 500
	costPerDist         uint64 = 1
)

var (
	comma    = []byte(",")
	origin   = newLocation(Home, 0, 0)
	homeLoad = NewLoad(0, origin, origin, true)
)

type locationType string

const (
	// Pickup indicates that the location is a pickup point
	Pickup locationType = "P"
	// Dropoff indicates that the location is a dropoff point
	Dropoff locationType = "D"
	// Home indicates the origin, which is neither a pickup or a dropoff point
	Home locationType = "H"
)
