/*
Write a function that uses variadic parameters to return the largest number provider
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("max(1, 2, 3) is ", max(1,2,3))
	fmt.Println("max(-1, -2, -3) is ", max(-1,-2,-3))
}

func max(vals ...float64) float64 {
	var retval float64 = math.Inf(-1)
	for _, 	val := range vals {
		if val > retval {
			retval = val
		}
	}
	return retval
}
