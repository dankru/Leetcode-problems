package main

import "fmt"

func main() {
	numbers := []int{2, 5, 5, 11}
	target := 10
	fmt.Println(twoSum(numbers, target))
}

// We will just check every possible sum until we find target
func twoSum(numbers []int, target int) [2]int {
	offset := 1
	for i, _ := range numbers {
		numbersOffset := numbers[offset:]
		for j, _ := range numbersOffset {
			if (numbers[i] + numbers[j+offset]) == target {
				return [2]int{i, j + offset}
			}
		}
		offset += 1
	}
	return [2]int{0, 0}
}
