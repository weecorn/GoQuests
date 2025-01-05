package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Point represents a point in 2D space
type Point struct {
	X, Y float64
}

// Distance calculates the Euclidean distance between two points
func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

// Line represents a line in 2D space defined by two points
type Line struct {
	P1, P2 Point
}

// Slope calculates the slope of a line
func (l Line) Slope() float64 {
	if l.P1.X == l.P2.X {
		return math.Inf(1) // Vertical line
	}
	return (l.P2.Y - l.P1.Y) / (l.P2.X - l.P1.X)
}

// Intercept calculates the y-intercept of a line
func (l Line) Intercept() float64 {
	return l.P1.Y - l.Slope()*l.P1.X
}

// Intersection calculates the intersection point of two lines
func Intersection(l1, l2 Line) (Point, error) {
	if l1.Slope() == l2.Slope() {
		return Point{}, fmt.Errorf("lines are parallel")
	}

	x := (l2.Intercept() - l1.Intercept()) / (l1.Slope() - l2.Slope())
	y := l1.Slope()*x + l1.Intercept()

	return Point{X: x, Y: y}, nil
}

// Circle represents a circle in 2D space
type Circle struct {
	Center Point
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Circumference calculates the circumference of a circle
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// GenerateRandomPoints generates a slice of n random points within a given rectangle
func GenerateRandomPoints(n int, rectX, rectY, rectWidth, rectHeight float64) []Point {
	rand.Seed(time.Now().UnixNano())
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = Point{
			X: rectX + rand.Float64()*rectWidth,
			Y: rectY + rand.Float64()*rectHeight,
		}
	}
	return points
}

func main() {
	// Example usage:
	p1 := Point{1, 2}
	p2 := Point{4, 6}
	fmt.Println("Distance between p1 and p2:", Distance(p1, p2))

	line1 := Line{p1, p2}
	fmt.Println("Slope of line1:", line1.Slope())
	fmt.Println("Intercept of line1:", line1.Intercept())

	line2 := Line{Point{2, 1}, Point{5, 5}}
	intersection, err := Intersection(line1, line2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Intersection of line1 and line2:", intersection)
	}

	circle := Circle{Point{0, 0}, 5}
	fmt.Println("Area of circle:", circle.Area())
	fmt.Println("Circumference of circle:", circle.Circumference())

	points := GenerateRandomPoints(10, 0, 0, 100, 100)
	fmt.Println("Random points:", points)
}
