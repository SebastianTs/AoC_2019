package main

import (
	"fmt"
)

const (
	start  = 168630
	end    = 718098
	digits = 6
)

func main() {
	count, part2Count := 0, 0
	cur := [digits]int{}
	for i := start; i < end; i++ {
		cur = intToArray(i)
		if meetsCriteria(cur) {
			count++
		}
		if meetsCriteriaExactDouble(cur) {
			part2Count++
		}
	}
	fmt.Printf("Solution 1:\t %d\nSolution 2:\t%d\n", count, part2Count)
}

func meetsCriteria(lst [digits]int) bool {
	hasDouble := false
	notDecrasing := false
	for i := 0; i < len(lst)-1; i++ {
		if lst[i] == lst[i+1] {
			hasDouble = true
		}
		notDecrasing = lst[i] <= lst[i+1]
		if !notDecrasing {
			break
		}
	}
	return hasDouble && notDecrasing
}

func meetsCriteriaExactDouble(lst [digits]int) bool {
	hasDouble := false
	notDecrasing := false
	count := map[int]int{}
	doubles := []int{}
	for i := 0; i < len(lst); i++ {
		count[lst[i]]++
	}
	for i := 0; i < len(lst)-1; i++ {
		if lst[i] == lst[i+1] {
			doubles = append(doubles, lst[i])
		}
		notDecrasing = lst[i] <= lst[i+1]
		if !notDecrasing {
			break
		}
	}
	for _, val := range doubles {
		if count[val] == 2 {
			hasDouble = true
			break
		}
	}
	return hasDouble && notDecrasing
}
func intToArray(number int) (out [digits]int) {
	for i := 0; i < digits; i++ {
		out[digits-i-1] = number % 10
		number /= 10
	}
	return out
}
