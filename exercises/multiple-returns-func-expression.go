/*
Write a function 'halfer' that returns a function takes an integer & returns 2 values
  -- First return value should be parameter divided by 2
  -- Second return value should be true if parameter is even and false otherwise

  Ex. half(1) => 0 false
  Ex. half(2) => 1 true
 */
package main

import "fmt"
func main() {
	halferFunc := halfer()
	retstat, retval := halferFunc(1)

	fmt.Println("halferFunc(1) - ", retstat, retval)

	retstat, retval = halferFunc(2)
	fmt.Println("halferFunc(2) - ", retstat, retval)

}

/*
halfer is a func that
	returns a func that takes an int & returns:
		- half the int and
		- true if the int is even and false otherwise
 */
func halfer() func(i int) (int, bool) {
	return func(i int) (int, bool) {
		return i / 2, i%2 == 0
	}
}
