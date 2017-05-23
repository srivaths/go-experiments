/*
 Project Euler - Problem #3
 The prime factors of 13195 are 5, 7, 13 and 29.
 What is the largest prime factor of the number 600851475143 ?
 */
package main

import (
	"fmt"
	"os"
	"sync"
	"math"
)

func main() {
	var targetNumber, foo int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &targetNumber)
	if foo = int(math.Abs(float64(targetNumber))); foo < 2 {
		fmt.Println("Bye")
		os.Exit(0)
	}
	findFactor(foo)
	fmt.Printf("(%d) Largest prime factor of %d is %d\n", threadcount, targetNumber, largestFactor)
}

var wg sync.WaitGroup
var largestFactor int
var threadcount int
var mutex sync.Mutex
func findFactor(targetNumber int) {
	// A factor of a given number cannot be greater than half of that number.
	var largestPossibleFactor = targetNumber / 2

	for i := largestPossibleFactor; i >= 2; i-- {
		wg.Add(1)
		threadcount++
		// Test only if i is a prime number
		go fullyDivisible(targetNumber, i)
	}
	wg.Wait()
}

/*
Tests if targetNumber is fully divisble by factor
 */
func fullyDivisible(targetNumber, factor int) {
	if targetNumber%factor == 0 && factor%2 != 0 && isPrime(factor) {
		// Set the largest factor only if this is larger
		// It *could* be smaller due to multi-threading.
		mutex.Lock()
		if factor > largestFactor {
			largestFactor = factor
		}
		mutex.Unlock()
	}
	wg.Done()
}

/*
Identifies if numberToTest is a prime number.
 */
func isPrime(numberToTest int) bool {
	for j := 2; j <= numberToTest/2; j++ {
		if largestFactor > 2 {
			return false
		}
		if numberToTest % j == 0 {
			return false
		}
	}
	return true
}
