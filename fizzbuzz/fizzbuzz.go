package main

import (
	"fmt"
)

/*
http://wiki.c2.com/?FizzBuzzTest
"Write a program that prints the numbers from 1 to 100.
But for multiples of three print “Fizz” instead of the number and
for the multiples of five print “Buzz”.
For numbers which are multiples of both three and five print “FizzBuzz”."
 */
func main() {
	var x string
	for i := 1; i < 101; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				x = "fizzbuzz"
			} else {
				x = "fizz"
			}
		} else if i % 5 == 0 {
			x = "buzz"
		} else {
			x = ""
		}
		fmt.Printf("%3d: %s\n", i, x)
	}
}
