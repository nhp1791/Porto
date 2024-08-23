package models

import (
	"math/rand"
	"strconv"
	"strings"
)

// Driver represents an individual Driver as part of a Driver stable
// which is a group of drivers used to complete a load set
type Driver struct {
	network        *LoadSet
	shiftSqDist    uint64
	load           *Load
	completedLoads []*Load
}

func newDriver(network *LoadSet) *Driver {
	return &Driver{
		network:        network,
		load:           homeLoad,
		completedLoads: []*Load{},
	}
}

// FindNearestPickup is called in the solution algorithm to determine whether or not the
// driver is able to find a pickup location for which it could complete the delivery and return home.
// If this is not possible, this function returns false
func (d *Driver) FindNearestPickup(choice int) bool {
	// Find the current node number of the driver.  Because of the way things are written,
	// the driver is at the dropoff point of this load.
	current := d.load.number
	size := d.network.size
	// Create a new neighborhood that will hold the nearest uncompleted neighbors, up to
	// a maximum of MaxNearestNeighbors.  Note that nearest is measured by the distance
	// between the current load's dropoff point and the neighbor's pickup point
	neighbors := newNeighborhood()

	// Cycle through all loads
	for i := 0; i < size; i++ {
		// Do not include the current load in the list of neighbors
		if i == current {
			continue
		}

		load := d.network.LoadMap[i]

		// Exclude completed loads
		if load.complete {
			continue
		}

		// Lookup the distance from the current dropoff point
		// to the neighbors pickup point
		dist := d.network.Matrix[current][i]
		// Insert that distance, if it's one of the nearest neighbors
		neighbors.insert(dist, load)
	}

	// Remove any nil neighbors (happens when there are not enough non-completed)
	// loads to fill a slice of length MaxNearestNeighbors
	neighbors.prune()

	// See how many nearest neighbors there are
	numNeighbors := len(neighbors.neighbors)

	// If there are no non-completed neighbors, we're done
	if numNeighbors == 0 {
		return true
	}

	// If our trial iteration is bigger than the remaining number of neighbors,
	// ensure we don't index out of bounds
	if choice > numNeighbors {
		choice = numNeighbors
	} else if choice < 0 {
		// For negative choices, choose a random nearest neighbor
		choice = rand.Intn(numNeighbors)
	}

	// Check each neighbor, starting at the indicated value (choice)
	// and find the first one that allows you to complete that load
	// and still be able to get home if necessary.
	for i := 0; i < numNeighbors; i++ {
		neighbor := neighbors.neighbors[(choice+i)%numNeighbors]
		if ok := d.testNeighbor(neighbor); ok {
			d.driveNeighbor(neighbor)
			return false
		}
	}

	// If there were neighbors, but we cannot complete
	// any of them, let the calling algorithm know that this
	// driver is done and needs to be sent home.
	return true
}

func (d *Driver) getCompletedLoadList() []*Load {
	return d.completedLoads
}

func (d *Driver) completedLoadString() string {
	c := make([]string, len(d.completedLoads))
	for i, l := range d.completedLoads {
		c[i] = strconv.Itoa(l.number)
	}

	return strings.Join(c, ",")
}

// ReturnHome moves a driver from the dropoff location of the
// current load back to the origin
func (d *Driver) ReturnHome() {
	d.shiftSqDist += d.network.Matrix[d.load.number][0]
	d.load = homeLoad
}

// testNeighbor simply checks to see if the driver could
// move from the current dropoff location to the neighboring
// pickup location, deliver that load, and return home without
// exceeding the shift limit.
func (d *Driver) testNeighbor(n *neighbor) bool {
	return d.shiftSqDist+
		n.dist+
		d.network.Matrix[n.load.number][n.load.number]+
		d.network.Matrix[0][n.load.number] < maxDriverSqDistance
}

// driveNeighbor actually executes a movement from a point to a neighboring
// pickup point, delivers the load, and sets the driver location to the
// new dropoff point.
func (d *Driver) driveNeighbor(n *neighbor) {
	d.shiftSqDist += n.dist + d.network.Matrix[n.load.number][n.load.number]
	d.load = n.load
	n.load.complete = true
	d.completedLoads = append(d.completedLoads, n.load)
}
