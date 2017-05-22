/*
 Project Euler - Problem #3
 The prime factors of 13195 are 5, 7, 13 and 29.
 What is the largest prime factor of the number 600851475143 ?
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	var targetNumber int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &targetNumber)
	if targetNumber < 2 {
		fmt.Println("Very funny.")
		os.Exit(0)
	}
	fmt.Printf("Largest prime factor of %d is %d\n", targetNumber, findFactor(targetNumber))
}

func findFactor(targetNumber int) int {
	// A factor of a given number cannot be greater than half of that number.
	var largestFactor = targetNumber / 2
	//fmt.Println("Largest possible prime of ", targetNumber, " is ", largestFactor)
	for i := largestFactor; i >= 2; i-- {
		// Test only if i is a prime number
		if targetNumber%i == 0 && i%2 != 0 && isPrime(i) {
				return i
		}
	}
	return targetNumber
}

func isPrime(numberToTest int) bool {
	for j := 2; j <= numberToTest/2; j++ {
		if numberToTest % j == 0 {
			//fmt.Println(numberToTest, " is not a prime number")
			return false
		}
	}
	//fmt.Println(numberToTest, " is a prime number")
	return true
}
