package main

func isPalindrome(x int32) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	figures := []int32{}

	for x != 0 {
		figures = append(figures, x%10)
		x /= 10
	}

	for i := 0; i < len(figures)/2; i++ {
		if figures[i] != figures[len(figures)-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeBest() {

}
