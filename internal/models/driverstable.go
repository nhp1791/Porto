package models

import "math"

// DriverStable is a set of drivers (picture the TV show Taxi) that can be used
// to completely deliver a load set
type DriverStable struct {
	loadset           *LoadSet
	dispatchedDrivers []*Driver
	cost              uint64
}

// NewDriverStable is a factory function for creating a new stable of drivers
func NewDriverStable(loadset *LoadSet) *DriverStable {
	return &DriverStable{
		loadset:           loadset,
		dispatchedDrivers: []*Driver{},
	}
}

// DispatchNewDriver creates a new driver whenever a previous driver has reached
// its limit
func (s *DriverStable) DispatchNewDriver() *Driver {
	driver := newDriver(s.loadset)
	s.dispatchedDrivers = append(s.dispatchedDrivers, driver)
	return driver
}

// CalculateCost creates a close estimate of the cost of a particular solution.
// Note that its an estimate because distances are calculated using rounded
// integer locations rather than the full floating values given in the problem.
// Statistically, the error should reduce toward zero as the problem set gets bigger.
func (s *DriverStable) CalculateCost() uint64 {
	drivers := uint64(len(s.dispatchedDrivers))

	var totalSqDist uint64
	for _, d := range s.dispatchedDrivers {
		totalSqDist += d.shiftSqDist
	}
	dist := uint64(math.Round(math.Sqrt(float64(totalSqDist))))

	return costPerDriver*drivers + costPerDist*dist
}

// Solution returns a slice of the strings representing the routes of each driver
func (s *DriverStable) Solution() []string {
	loadStrings := make([]string, len(s.dispatchedDrivers))
	for i, d := range s.dispatchedDrivers {
		loadStrings[i] = d.completedLoadString()
	}
	return loadStrings
}

// Size returns the total number of loads in a solution, along with the total number of
// unique loads.  These should always be the same.
func (s *DriverStable) Size() (total int, unique int) {
	set := make(map[int]bool)
	for _, d := range s.dispatchedDrivers {
		driverSet := d.getCompletedLoadList()
		total += len(driverSet)
		for _, l := range driverSet {
			set[l.number] = true
		}
	}
	return total, len(set)
}
