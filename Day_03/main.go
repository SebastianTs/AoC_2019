package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)
type move struct{
	dir rune
	dist int
}

type wire struct{
	ms []move
}

type tuple struct{
	x int
	y int
}

func main(){
	ns, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(ns))
	findCollision(ns)
}

func findCollision(wires []wire) int{
	field := make(map[tuple]int)
	for i, wire := range wires{
		for _, move := range wire.ms{
			if move.dir == 'R' {
				fmt.Println(string(move.dir),i)
			}
		}
	}
	return 0
}

func readInput(filename string) (wires []wire, err error) {
	wires = make([]wire,0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return wires, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		scntmp := bufio.NewScanner(strings.NewReader(line))
		onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			for i := 0; i < len(data); i++ {
				if data[i] == ',' {
					return i + 1, data[:i], nil
				}
			}
			if !atEOF {
				return 0, nil, nil
			}
			return 0, data, bufio.ErrFinalToken
		}
		scntmp.Split(onComma)
		ms := make([]move, 0)	
		for scntmp.Scan() {
			var dist int
			cur := scntmp.Text()
			fmt.Sscanf(cur[1:],"%d", &dist)
			ms = append(ms, move{rune(cur[0]),dist})
		}
		wires = append(wires, wire{ms})
	}
	return wires, nil
}