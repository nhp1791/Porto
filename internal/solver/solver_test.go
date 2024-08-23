package solver

import (
	"sched/internal/reader"
	"testing"
)

func TestSolution(t *testing.T) {
	loadset := reader.CreateLoadSet("./testfiles/problem.txt")
	if loadset == nil {
		t.Fatal("could not read problem file")
	}

	solution := SolveLoadSet(loadset, false)

	actualSolution := []string{
		"88,36,13,176,101,183,147,71,15,28,173,163,168,38,153,86,130,17,184,121,115,1,44,62,125,18,145,104,2,19,133,165,185,177,200,45,193,81,198,92,102,85,152,14,109",
		"117,129,100,106,136,154,66,21,151,181,174,144,47,75,105,195,156,192,166,124,79,134,196,122,98,108,77,178,53,60,12,119,116,78,67,49,83",
		"160,91,73,68,48,146,143,167,161,35,137,23,175,52,84,59,141,194,89,126,61,82,159,103,58,170,54,158,74,30,37,69,7,164",
		"6,179,190,72,162,197,46,11,138,41,65,191,87,33,27,107,140,169,149,26,114,95,40,150,97,155,34,31,172,157,70,96,93",
		"120,112,135,180,139,16,187,182,110,113,22,5,20,142,50,32,25,171,51,132,94,128,90,3,39,56,148,123,189,199,76",
		"131,186,99,111,57,63,80,42,43,9,10,24,4,188,118,127,55,64,29,8",
	}

	if len(solution) != len(actualSolution) {
		t.Fatal("wrong size solution")
	}

	for i, l := range solution {
		if l != actualSolution[i] {
			t.Fatalf("wrong solution.  wanted='%s', got='%s'", actualSolution[i], l)
		}
	}
}
