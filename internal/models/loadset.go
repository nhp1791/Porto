package models

// LoadSet is the collection of loads that make up a full problem to be solved
type LoadSet struct {
	size int
	// LoadMap maps the load number to the load struct
	LoadMap map[int]*Load
	Matrix  [][]uint64
}

// NewLoadSet is a factory function for creating a new LoadSet.
// Notice that the origin is added to each new LoadSet
func NewLoadSet() *LoadSet {
	loadset := &LoadSet{
		LoadMap: make(map[int]*Load),
	}
	loadset.AddLoad(homeLoad)
	return loadset
}

// Clone creates a new LoadSet that is identical
// to the provided one, except that all loads
// are set to not complete.
func (l *LoadSet) Clone() *LoadSet {
	n := NewLoadSet()
	n.size = l.size
	for i := 1; i < n.size; i++ {
		n.LoadMap[i] = l.LoadMap[i].clone()
	}
	n.Matrix = l.Matrix

	return n
}

// AddLoad simply adds a load to the LoadMap
func (l *LoadSet) AddLoad(load *Load) {
	l.LoadMap[load.number] = load
}

// FormDistanceMatrix is called once all Loads have been added to the LoadSet,
// This method creates the matrix that calculates the distance
// from Dropoff of the row number to the Pickup of the column number,
// where the row and column numbers are the same as the load number
func (l *LoadSet) FormDistanceMatrix() {
	size := len(l.LoadMap)
	l.size = size
	matrix := make([][]uint64, size)
	for k, v := range l.LoadMap {
		row := make([]uint64, size)
		for i := 0; i < size; i++ {
			row[i] = v.Dropoff.sqDistance(l.LoadMap[i].Pickup)
		}
		matrix[k] = row
	}
	l.Matrix = matrix
}

// IsFinished just checks to see if an uncompleted load still exists
func (l *LoadSet) IsFinished() bool {
	for _, v := range l.LoadMap {
		if !v.complete {
			return false
		}
	}
	return true
}

// Size returns the total number of loads to be solved for.
func (l *LoadSet) Size() int {
	// Because the origin in added to the LoadMap, but is
	// not an actual load, one must subtracted from the LoadMap size
	return len(l.LoadMap) - 1
}
