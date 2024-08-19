// Package reader implements the ability to read problem files of the form
//
// loadNumber pickup dropoff
// 1 (-9.100071078494038,-48.89301103772511) (-116.78442279683607,76.80147820713637)
// 2 (73.38933871575719,-86.93443314676254) (-57.594533352956425,28.662926099543245)\
// ...
//
// # Returning a LoadSet struct
//
// CAVEAT: The problem file is assumed to label the points consecutively starting at 1.
// If this is not the case, pre-processing of the file is needed.
package reader

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sched/internal/models"
	"strconv"
)

// CreateLoadSet reads a problem file and returns a LoadSet struct
// consisting of all the loads defined in the file, along with a load
// representing the origin, labeled as the 0 load.
func CreateLoadSet(filename string) *models.LoadSet {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		println(err.Error())
		return nil
	}
	defer f.Close()

	// Create the loadset to return
	loadset := models.NewLoadSet()

	// Process the file line-by-line
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		val := sc.Bytes()
		loadNumber, pickup, dropoff := processLine(val)
		if loadNumber == -1 {
			continue
		}
		if pickup == nil {
			return nil
		}

		load := models.NewLoad(loadNumber, pickup, dropoff, false)
		loadset.AddLoad(load)
	}

	loadset.FormDistanceMatrix()

	return loadset
}

// For each line, extract the load number and the pickup and dropoff locations.  If there is a
// problem, print the problem and return nil locations so that the reader can stop.
func processLine(val []byte) (loadNumber int, pickup *models.Location, dropoff *models.Location) {
	vals := bytes.Split(val, []byte(" "))

	if len(vals) != 3 {
		_, _ = fmt.Printf("Line '%s' did not have three fields.  Skipping", val)
		return 0, nil, nil
	}

	ln := string(vals[0])
	loadNum, err := strconv.ParseInt(ln, 10, 0)
	if err != nil {
		if string(ln) != "loadNumber" {
			_, _ = fmt.Printf("Line '%s' did not have an integer to start it.  Skipping", val)
			return 0, nil, nil
		}
		return -1, nil, nil
	}

	loadNumber = int(loadNum)

	pickup = models.FormLocation(vals[1], models.Pickup)
	if pickup == nil {
		_, _ = fmt.Printf("Line '%s' did not have a location as its second element.  Skipping", val)
		return 0, nil, nil
	}

	dropoff = models.FormLocation(vals[2], models.Dropoff)
	if dropoff == nil {
		_, _ = fmt.Printf("Line '%s' did not have a location as its third element.  Skipping", val)
		return 0, nil, nil
	}

	return loadNumber, pickup, dropoff
}
