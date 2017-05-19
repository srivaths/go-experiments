/*
Write a function 'half' that takes an integer & returns 2 values
  -- First return value should be parameter divided by 2
  -- Second return value should be true if parameter is even and false otherwise

  Ex. half(1) => 0 false
  Ex. half(2) => 1 true
 */
package main

import "fmt"
func main() {
	retstat, retval := half(1)

  fmt.Println("half(1) - ", retstat, retval)

	retstat, retval = half(2)
	fmt.Println("half(2) - ", retstat, retval)

}
func half(i int) (int, bool) {
	return i/2, i%2==0
}
