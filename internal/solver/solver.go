// Package solver contains an algorithm framework
// for finding a good solution to a loading problem.
// Note that the majority of work in this algorithm is
// actually performed in the FindNearestPickup method of
// the driver struct
package solver

import (
	"fmt"
	"math"
	"sched/internal/models"
)

// SolveLoadSet is called to produce a solution to the loading
// problem and to print out the solution in the form
//
// 1,6,5
// 2,3
// 4
//
// where each line represents the route of an individual driver
func SolveLoadSet(loadset *models.LoadSet, debug bool) []string {
	var minCost uint64 = math.MaxUint64
	var bestStable *models.DriverStable

	// We're going to try 2 * models.MaxNearestNeighbors paths through the system.
	// The negative values will effectively be a set of Monte Carlo experiments, varying the
	// choice of next load (within a set of size models.MaxNearestNeighbors) randomly,
	// while the positive values will choose that neighbor (i.e. i = 0 chooses the nearest neighbor
	// every time, i=1 chooses the next nearest neighbor every time, etc.) Note that nearest
	// neighbor in this sense is defined as the non-completed load who's pickup point is closest
	// to the current dropoff point.
	for i := -models.MaxNearestNeighbors; i < models.MaxNearestNeighbors; i++ {
		// Clone the loadset so that nodes are initially marked as not completed
		ls := loadset.Clone()

		// Create a stable of drivers
		stable := models.NewDriverStable(ls)
		// Get a new driver to start the solution
		driver := stable.DispatchNewDriver()
		// While there are loads that have not been completed, continue the algorithm
		for !ls.IsFinished() {
			// Find out if the driver can complete another load
			finished := driver.FindNearestPickup(i)
			// If the driver was unable to pickup another load, send it
			// home and get a new driver
			if finished {
				driver.ReturnHome()
				driver = stable.DispatchNewDriver()
				continue
			}
		}

		// When all loads have been completed, the last driver is at the dropoff
		// location of the final load, and needs to get back home
		driver.ReturnHome()

		// Calculate the cost for this solution
		cost := stable.CalculateCost()

		// Keep track of the minimum cost solution
		if cost < minCost {
			minCost = cost
			bestStable = stable
		}
	}

	// If no solution was found, say so.
	if bestStable == nil {
		println("No solution could be found")
		return []string{}
	}

	if debug {
		size, uniqueSize := bestStable.Size()
		_, _ = fmt.Printf("Size of solution set: %d\n", size)
		_, _ = fmt.Printf("Number of unique solutions in solution set: %d\n", uniqueSize)
		println()
	}

	// Print the loads for each driver
	return bestStable.Solution()
}
