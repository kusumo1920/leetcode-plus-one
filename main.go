package main

import "fmt"

func main() {
	digits := []int{9, 9, 9}
	output := plusOneSolution1(digits)
	traverseSlice(output)
}

func plusOneSolution1(digits []int) []int {
	result := make([]int, len(digits)+1)

	tempStore := -1
	sum := digits[len(digits)-1] + 1
	if sum/10 > 0 {
		tempStore = sum / 10
	}
	result[len(digits)] = sum % 10

	for i := len(digits) - 2; i >= 0; i-- {
		currentDigit := digits[i]
		if tempStore != -1 {
			currentDigit += tempStore
			tempStore = -1
		}

		if currentDigit/10 > 0 {
			tempStore = currentDigit / 10
		}

		result[i+1] = currentDigit % 10
	}

	if tempStore != -1 {
		result[0] = tempStore
		tempStore = -1
	} else {
		return result[1:]
	}

	return result
}

func traverseSlice(nums []int) {
	fmt.Print("[")
	for i, num := range nums {
		fmt.Print(num)
		if i+1 < len(nums) {
			fmt.Print(" ")
		}
	}
	fmt.Print("]")
}
