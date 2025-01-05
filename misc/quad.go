package main

import (
	"fmt"
	"math"
)

// This program calculates various properties of a quadratic equation
// in the form of ax^2 + bx + c = 0

// quadraticRoots calculates the roots of a quadratic equation
func quadraticRoots(a, b, c float64) (float64, float64, error) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, fmt.Errorf("discriminant is negative, no real roots")
	}

	root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	root2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return root1, root2, nil
}

// vertex calculates the vertex of the parabola represented by the quadratic equation
func vertex(a, b, c float64) (float64, float64) {
	x := -b / (2 * a)
	y := a*x*x + b*x + c
	return x, y
}

// isPerfectSquare checks if a number is a perfect square
func isPerfectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

// findPythagoreanTriplets finds all Pythagorean triplets within a given limit
func findPythagoreanTriplets(limit int) [][]int {
	var triplets [][]int
	for a := 1; a <= limit; a++ {
		for b := a; b <= limit; b++ {
			cSquared := a*a + b*b
			if isPerfectSquare(cSquared) && int(math.Sqrt(float64(cSquared))) <= limit {
				triplets = append(triplets, []int{a, b, int(math.Sqrt(float64(cSquared)))})
			}
		}
	}
	return triplets
}

func main() {
	// Calculate roots of a quadratic equation
	a, b, c := 1.0, -3.0, 2.0
	root1, root2, err := quadraticRoots(a, b, c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Roots of the quadratic equation: %.2f, %.2f\n", root1, root2)
	}

	// Calculate vertex of the parabola
	vx, vy := vertex(a, b, c)
	fmt.Printf("Vertex of the parabola: (%.2f, %.2f)\n", vx, vy)

	// Find Pythagorean triplets
	limit := 20
	triplets := findPythagoreanTriplets(limit)
	fmt.Println("Pythagorean triplets within limit", limit, ":")
	for _, triplet := range triplets {
		fmt.Println(triplet)
	}
}
