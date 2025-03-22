package main

// We will just check every possible sum until we find target
func twoSum(numbers []int, target int) []int {
	offset := 1
	for i, _ := range numbers {
		numbersOffset := numbers[offset:]
		for j, _ := range numbersOffset {
			if (numbers[i] + numbers[j+offset]) == target {
				return []int{i, j + offset}
			}
		}
		offset += 1
	}
	return []int{}
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
