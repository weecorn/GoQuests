package main

import (
	"fmt"
	"math"
)

// This program explores various mathematical concepts and algorithms

// factorial calculates the factorial of a non-negative integer
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// fibonacci calculates the nth Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// gcd calculates the greatest common divisor of two integers using Euclid's algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the least common multiple of two integers
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	// Calculate factorial
	n := 5
	fmt.Println("Factorial of", n, ":", factorial(n))

	// Check if a number is prime
	num := 17
	if isPrime(num) {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// Calculate Fibonacci number
	fibNum := 10
	fmt.Println("Fibonacci number at position", fibNum, ":", fibonacci(fibNum))

	// Calculate GCD and LCM
	a, b := 24, 36
	fmt.Println("GCD of", a, "and", b, ":", gcd(a, b))
	fmt.Println("LCM of", a, "and", b, ":", lcm(a, b))
}
