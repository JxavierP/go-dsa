package main

import (
	"fmt"
	"time"
)

func computeSum(numbers []int) int {
	// Initialize a variable to hold the sum of the numbers
	sum := 0
	// Iterate through the array and add each number to the sum
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func computeAverage(numbers []int) float64 {
	// If the array is empty, return 0 to avoid division by zero
	if len(numbers) == 0 {
		return 0 
	}
	// Call computeSum to get the sum of the numbers
	sum := computeSum(numbers)
	// Calculate the average by dividing the sum by the number of elements
	return float64(sum) / float64(len(numbers))
}

func doubleValues(numbers []int) []int {
	// Create a new array to hold the doubled values
	newArray := []int{}

	// Iterate through the original array and append each value twice to the new array
	for _, number := range numbers {
		newArray = append(newArray, number)
		newArray = append(newArray, number)
	}
	return newArray
}

func uniqueString(arr []string) []string {
	// Create a map to track if a string is already present
	uniqueMap := make(map[string]bool)
	// Create a slice to store unique strings
	uniqueSlice := []string{}
	
	// Iterate through the array and add unique strings to the map and slice
	for _, string := range arr {
		// If the string is not in the map, add it to both the map and the slice
		if _, exists := uniqueMap[string]; !exists {
			uniqueMap[string] = true
			uniqueSlice = append(uniqueSlice, string)
		}
	}
	// Return the slice containing unique strings
	return uniqueSlice
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	startTime := time.Now()
	sum := computeSum(numbers)
	elapsedTime := time.Since(startTime)
	fmt.Printf("THe sum of the numbers is: %d and it took %s to run\n", sum, elapsedTime)
	fmt.Println("This function time complexity is 0(n) because it depends on the number of elements in the array.")

	average := computeAverage(numbers)
	fmt.Printf("The average of the numbers is: %.2f\n", average)
	fmt.Println("This function time complexity is also 0(n) because it also depends on the number of elements in the array.")

	doubledNumbers := doubleValues(numbers)
	fmt.Printf("The doubled values are: %v\n", doubledNumbers)
	fmt.Println("This function space complexity is 0(n) because it depends on the number of elements in the array, as it creates a new array with double the size.")

	strings := []string{"apple", "banana", "apple", "orange", "banana"}
	uniqueStrings := uniqueString(strings)
	fmt.Printf("The unique strings are: %v\n", uniqueStrings)
}
