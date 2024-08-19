package reader

import "testing"

func TestBadInput(t *testing.T) {
	loadset := CreateLoadSet("non-existent-file")
	if loadset != nil {
		t.Fatal("failed to properly respond to non-existent file")
	}
}

func TestEmptySet(t *testing.T) {
	loadset := CreateLoadSet("./testfiles/empty.txt")
	if loadset == nil {
		t.Fatal("should have read existent, empty file")
	}
	if len(loadset.LoadMap) != 1 {
		t.Fatal("should have gotten a LoadMap with just the origin in the loadset")
	}
	o, ok := loadset.LoadMap[0]
	if !ok {
		t.Fatal("should have gotten the origin labelled with 0 in the loadset")
	}
	if o.Dropoff.X != 0 || o.Dropoff.Y != 0 || o.Pickup.X != 0 || o.Pickup.Y != 0 {
		t.Fatal("improper origin in loadset")
	}
}

func TestSmallSet(t *testing.T) {
	loadset := CreateLoadSet("./testfiles/single.txt")
	if loadset == nil {
		t.Fatal("should have read existent file")
	}
	if len(loadset.LoadMap) != 2 {
		t.Fatal("should have gotten a LoadMap with the origin and one other point in the loadset")
	}
	o, ok := loadset.LoadMap[0]
	if !ok {
		t.Fatal("should have gotten the origin labelled with 0 in the loadset")
	}
	if o.Dropoff.X != 0 || o.Dropoff.Y != 0 || o.Pickup.X != 0 || o.Pickup.Y != 0 {
		t.Fatal("improper origin in loadset")
	}
	o, ok = loadset.LoadMap[1]
	if !ok {
		t.Fatal("should have gotten a point labelled with 1 in the loadset")
	}
	if o.Dropoff.X != -117 || o.Dropoff.Y != 77 || o.Pickup.X != -9 || o.Pickup.Y != -49 {
		t.Fatal("improper read of point in loadset")
	}
}

func TestNoLoadNumberError(t *testing.T) {
	loadset := CreateLoadSet("./testfiles/no_load_number.txt")
	if loadset != nil {
		t.Fatal("failed to properly read file containing a missing load number")
	}
}

func TestBadLocationError(t *testing.T) {
	loadset := CreateLoadSet("./testfiles/bad_location.txt")
	if loadset != nil {
		t.Fatal("failed to properly read a file with a bad location")
	}
}

func TestMissingLocationError(t *testing.T) {
	loadset := CreateLoadSet("./testfiles/missing_location.txt")
	if loadset != nil {
		t.Fatal("failed to properly read a file with a missing location")
	}
}
