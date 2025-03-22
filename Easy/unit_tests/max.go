package main

import "fmt"

func main() {
	fmt.Println(Max([]int{5, 2, 1, 3, 4}))
}

func Max(numbers []int) int {
	var max int

	for _, v := range numbers {
		if max < v {
			max = v
		}
	}

	return max
}

func MaxIndex(numbers []int) int {
	var maxIndex int
	var max int

	for i, v := range numbers {
		if max < v {
			max = v
			maxIndex = i
		}
	}

	return maxIndex
}
