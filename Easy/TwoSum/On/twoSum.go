package main

import (
	"fmt"
)

func twoSumOn(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, num := range numbers {
		complement := target - num
		if index, found := m[complement]; found {
			return []int{index, i}
		}
		m[num] = i
	}
	return []int{}
}

// Собираем индексы, делаем lookup числа который дополнит текущую цифру до таргета
func twoSumAlternateOn(numbers []int, target int) []int {
	m := map[int]int{}
	// Заполняем map
	for i, v := range numbers {
		m[v] = i
	}
	// Проходим по слайсу в поиске таргета
	for i, v := range numbers {
		complement := target - v
		fmt.Println(complement)
		targetIndex, ok := m[complement]
		if ok && targetIndex != i {
			return []int{i, targetIndex}
		}
	}
	return []int{}
}
