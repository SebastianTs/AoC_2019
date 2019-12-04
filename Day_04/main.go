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
	c1, c2 := 0, 0
	cur := [digits]int{}
	for i := start; i < end; i++ {
		cur = intToArray(i)
		if meetsCriteria(cur) {
			c1++
		}
		if meetsCriteriaExactDouble(cur) {
			c2++
		}
	}
	fmt.Printf("Solution 1:\t%d\nSolution 2:\t%d\n", c1, c2)
}

func meetsCriteria(lst [digits]int) bool {
	hasDouble := false
	Decrasing := true
	for i := 0; i < len(lst)-1; i++ {
		if lst[i] == lst[i+1] {
			hasDouble = true
		}
		Decrasing = lst[i] > lst[i+1]
		if Decrasing {
			break
		}
	}
	return hasDouble && !Decrasing
}

func meetsCriteriaExactDouble(lst [digits]int) bool {
	hasDouble := false
	Decrasing := true
	count := map[int]int{}
	doubles := []int{}
	for i := 0; i < len(lst); i++ {
		count[lst[i]]++
	}
	for i := 0; i < len(lst)-1; i++ {
		if lst[i] == lst[i+1] {
			doubles = append(doubles, lst[i])
		}
		Decrasing = lst[i] > lst[i+1]
		if Decrasing {
			break
		}
	}
	for _, val := range doubles {
		if count[val] == 2 {
			hasDouble = true
			break
		}
	}
	return hasDouble && !Decrasing
}

func intToArray(number int) (out [digits]int) {
	for i := 0; i < digits; i++ {
		out[digits-i-1] = number % 10
		number /= 10
	}
	return out
}
