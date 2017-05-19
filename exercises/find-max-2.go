/*
Write a function that uses variadic parameters to return the largest number provider.
Run it against a list of values provided from slices.
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	fixture := [][]float64 {
		{1,2,3},
		{-1, 2, 3},
		{-11, -12, -13},
	}
	for _, sample := range fixture {
		fmt.Println("Max of ", sample, "is ", max2(sample...))
	}
}

func max2(vals ...float64) float64 {
	var retval float64 = math.Inf(-1)
	for _, 	val := range vals {
		retval = math.Max(retval, val)
	}
	return retval
}
