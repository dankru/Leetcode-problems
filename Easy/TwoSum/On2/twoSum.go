package main

import "fmt"

func main() {
	numbers := []int{2, 5, 5, 11}
	target := 10
	//fmt.Println(twoSum(numbers, target))
	fmt.Println(twoSumAlternate(numbers, target))
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

func twoSumAlternate(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{j, i}
			}
		}
	}
	return []int{}
}

// 5 5
