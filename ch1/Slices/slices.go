package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(size int) []int {
	var numbers = make([]int, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(100)
	}

	return numbers
}

func reverse(data []int) []int {
	var length int = len(data)
	var reversed = make([]int, length)
	copy(reversed, data)

	for i := 0; i < length/2; i++ {
		j := length - i - 1
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return reversed
}

func removeDuplicates(data []int) []int {
	var keys = make(map[int]bool)
	var unique = make([]int, 0)

	for _, item := range data {
		if _, ok := keys[item]; !ok {
			keys[item] = true
			unique = append(unique, item)
		}
	}
	return unique
}

func findMinMax(data []int) (int, int) {
	var min, max int = data[0], data[0]

	for _, item := range data {
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
	}
	return min, max
}

func main() {
	var data = generateRandomNumbers(20)
	reversed := reverse(data)
	fmt.Println("Generated: ", data)
	fmt.Println("Reversed: ", reversed)

	unique := removeDuplicates(data)
	fmt.Println("Unique: ", unique)

	min, max := findMinMax(data)
	fmt.Println("Minimum: ", min)
	fmt.Println("Maximum: ", max)
}
