package util

import (
	"strconv"
	"strings"
)

// All puzzle inputs stored as an array of UTF-8 strings.
var Inputs []string

// Parses each line of the input with the given parser function.
func ParseInput(input string, parser func(string) []string) [][]string {
	lines := strings.Split(input, "\n")
	r := make([][]string, 0)
	for _, line := range lines {
		fields := parser(line)
		if len(fields) == 0 {
			continue
		}
		r = append(r, fields)
	}

	return r
}

// Returns the input as a two-dimensional array of float64.
func InputFloats(input string, parser func(string) []string) [][]float64 {
	lines := ParseInput(input, parser)

	r := make([][]float64, len(lines))
	var err error
	for lineNo, fields := range lines {
		nums := make([]float64, len(fields))
		for i, f := range fields {
			nums[i], err = strconv.ParseFloat(f, 64)
			if err != nil {
				panic(err)
			}
		}
		r[lineNo] = nums
	}

	return r
}

// Returns the input as a two-dimensional array of int.
func InputInts(input string, parser func(string) []string) [][]int {
	lines := InputFloats(input, parser)

	r := make([][]int, len(lines))
	for lineNo, fields := range lines {
		nums := make([]int, len(fields))
		for i, f := range fields {
			nums[i] = int(f)
		}
		r[lineNo] = nums
	}

	return r
}

// CSVParser ...
func CSVParser(input string) []string {
	return strings.FieldsFunc(input, func(c rune) bool { return c == ',' })
}

func init() {
	Inputs = make([]string, 25)
	Inputs[0] = `1,2,3
4,5,6
7,8,9,10
`
	// As the inputs are released, store them right here inline. Simple.
	inputs[0] = `119965
69635
134375
71834
124313
109114
80935
146441
120287
85102
148451
69703
143836
75280
83963
108849
133032
109359
78119
104402
89156
116946
132008
131627
124358
56060
141515
75639
146945
95026
99256
57751
148607
100505
65002
78485
84473
112331
82177
111298
131964
125753
63970
77100
90922
119326
51747
104086
141344
54409
69642
70193
109730
73782
92049
90532
147093
62719
79829
142640
85242
128001
71403
75365
90146
147194
76903
68895
56817
142352
77843
64082
106953
115590
87224
58146
134018
127111
51996
134433
148768
103906
52848
108577
77646
95930
67333
98697
55870
78927
148519
68724
93076
73736
140291
121184
111768
71920
104822
87534
`
}
