package main

import (
	"testing"
)

var inputTest = []struct {
	input   string
	output1 int
	output2 int
}{
	{"aaoms", 3, 2},
	{"testingagain", 5, 7},
	{"HelloWorld", 3, 7},
}

func TestCount(t *testing.T) {
	for _, val := range inputTest {
		s, k := countHandler(val.input)
		if s != val.output1 || k != val.output2 {
			t.Errorf("countHandler(%s) => %d, %d Should be %d, %d /n", val.input, s, k, val.output1, val.output2)
		}
	}
}
