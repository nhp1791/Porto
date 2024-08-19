package models

import (
	"bytes"
	"math"
	"strconv"
)

// Location is a struct that holds a type (pickup, dropoff, home) and
// a cartesian coordinate that has been rounded off to the nearest integer
type Location struct {
	Type locationType
	X    int32
	Y    int32
}

func newLocation(t locationType, x int32, y int32) *Location {
	return &Location{
		Type: t,
		X:    x,
		Y:    y,
	}
}

// FormLocation creates a location from a string of the form
// (-9.100071078494038,-48.89301103772511)
func FormLocation(candidate []byte, t locationType) *Location {
	cand := bytes.TrimSpace(candidate)
	if cand[0] != lparen || cand[len(cand)-1] != rparen {
		return nil
	}

	coords := bytes.Split(cand[1:len(cand)-2], comma)
	if len(coords) != 2 {
		return nil
	}

	x, err := strconv.ParseFloat(string(coords[0]), 64)
	if err != nil {
		return nil
	}

	rx := int32(math.Round(x))

	y, err := strconv.ParseFloat(string(coords[1]), 64)
	if err != nil {
		return nil
	}

	ry := int32(math.Round(y))

	return newLocation(t, rx, ry)
}

func (l *Location) sqDistance(other *Location) uint64 {
	delX := other.X - l.X
	delY := other.Y - l.Y
	return uint64(delX*delX + delY*delY)
}
