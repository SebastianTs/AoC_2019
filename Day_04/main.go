package main

import (
	"fmt"
)

const(
	start = 168630
	end = 718098
	digits = 6
)

func main(){
	
	var candidates [][digits]int
	candidates = make([][digits]int, end-start)
	cur := [digits]int{}
	j:= 0
	for i:=start; i <end; i++ {
		cur = intToArray(i)
		candidates[j] = cur
		j++
	}
	fmt.Println(len(candidates))
	removeFromList(0, candidates)
	fmt.Println(candidates[0])
}

func hasDouble(lst [digits]int) bool{
	for i:=0; i < 10; i++{
		fmt.Println(i)
	} 
	return false

}

func removeFromList(i int, lst [][digits]int){
	lst[i] = lst[len(lst)-1]
	lst[len(lst)-1] = [digits]int{}
	lst = lst[:len(lst)-1]
}

func intToArray(number int) (out [digits]int){
	for i:= 0; i < digits; i++{
		out[digits-i-1] = number % 10
		number /= 10
	}
	return out
}