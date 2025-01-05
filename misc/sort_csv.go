package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Get the CSV string from command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <CSV string>")
		return
	}
	csvString := os.Args[1]

	// Split the CSV string into individual numbers
	stringNumbers := strings.Split(csvString, ",")
	var numbers []int

	// Convert the string numbers to integers
	for _, strNum := range stringNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return
		}
		numbers = append(numbers, num)
	}

	// Sort the numbers
	sort.Ints(numbers)

	// Print the sorted sequence
	fmt.Println("Sorted sequence:", numbers)

	// Convert the sorted numbers back to a CSV string
	var sortedCSV strings.Builder
	for i, num := range numbers {
		sortedCSV.WriteString(strconv.Itoa(num))
		if i < len(numbers)-1 {
			sortedCSV.WriteString(",")
		}
	}

	// Print the sorted sequence in CSV format
	fmt.Println("Sorted sequence (CSV):", sortedCSV.String())
}
