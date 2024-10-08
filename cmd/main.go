// This program creates a set of drivers/routes that is intended to be reasonably efficient
// in minimizing the number of drivers
//
// Usage:
//
//	./schedule -f /path/to/problem_file
package main

import (
	"flag"
	"fmt"
	"os"

	"sched/internal/reader"
	"sched/internal/solver"
)

func main() {
	// Get the filepath and quit if there is none
	var filepath string
	flag.StringVar(&filepath, "f", "", "The full path of the file containing the problem to be solved")

	var debug bool
	flag.BoolVar(&debug, "d", false, "Turns on debug printing")

	flag.Parse()

	if filepath == "" {
		_, _ = fmt.Println("Cannot proceed without problem file")
		os.Exit(1)
	}

	// Read the file and create a data structure holding the set of loads in the problem.
	// If there was a problem reading the file or creating the load set, exit.
	loadset := reader.CreateLoadSet(filepath)
	if loadset == nil {
		os.Exit(1)
	}

	if debug {
		println()
		_, _ = fmt.Printf("Number of loads requested: %d\n", loadset.Size())
	}

	// Find a reasonably efficient solution
	solution := solver.SolveLoadSet(loadset, debug)
	for _, s := range solution {
		_, _ = fmt.Printf("[%s]\n", s)
	}
}
