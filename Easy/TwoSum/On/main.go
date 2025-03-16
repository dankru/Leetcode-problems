package main

import "fmt"

func main() {
	numbers := []int{1, 5, 4, 4}
	target := 8
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) (int, int) {
	numbersList := map[int]int{}
	for i, v := range numbers {
		numbersList[i] = v

	}
	return 0, 0
}
