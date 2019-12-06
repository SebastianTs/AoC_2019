package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct{
	 A string
	 B string
}

type node struct{
	name string
	next []*node
	prev *node
}

func main() {
	ns, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	ns = []pair{
		{"COM","B"},{"B","C"},{"C","D"},{"D","E"},{"E","F"},
		{"B","G"},{"G","H"},{"D","I"},{"E","J"},{"J","K"},{"K","L"},
	}
	fmt.Println(calculateChecksum(ns))

}

func calculateChecksum(ns []pair) int{
	nodeMap := map[string]node{}
	for _, v := range ns{
		nodeMap[v.A] = node{name: v.A}
		nodeMap[v.B] = node{name: v.B}
	}
	for _, v := range ns{
		cur := nodeMap[v.B]
		n := nodeMap[v.A]
		cur.prev = &n
		nodeMap[v.B] = cur
		cur = nodeMap[v.A]
		n = nodeMap[v.B]
		cur.next = append(cur.next, &n)
		nodeMap[v.A] = cur
	}
	count := 0
	for _, v := range nodeMap{
		count += countNodes(&v)
	}
	return count
}

func countNodes(list *node) (c int){
	for list !=nil {
		list = list.prev
		c++
		fmt.Println(list)
	}
	return 
}




func readInput(filename string) (ns []pair, err error) {
	ns = make([]pair, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ns, err
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		a := line[:3]
		b := line[4:]
		ns = append(ns,pair{a,b})
	}
	return ns, nil
}