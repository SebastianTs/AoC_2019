package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type move struct {
	dir  rune
	dist int
}

type wire struct {
	ms []move
}

type tupel struct {
	x int
	y int
}

func main() {
	ns, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution to Part1\t%d\n", findCollision(ns))
	fmt.Printf("Solution to Part2\t%d\n", signalDelay(ns))
}

func findCollision(wires []wire) int {
	field := make(map[tupel][]bool)
	for i, wire := range wires {
		cur := tupel{0, 0}
		for _, move := range wire.ms {
			for j := 0; j < move.dist; j++ {
				switch move.dir {
				case 'R':
					cur.x++
				case 'L':
					cur.x--
				case 'U':
					cur.y++
				case 'D':
					cur.y--
				}
				if _, inUse := field[cur]; !inUse {
					field[cur] = make([]bool, len(wires))
				}
				field[cur][i] = true
			}
		}

	}
	distances := []int{}
	for k, v := range field {
		count := 0
		for _, entry := range v {
			if entry {
				count++
			}
		}
		if count > 1 {
			distances = append(distances, abs(k.x)+abs(k.y))
		}
	}
	return min(distances)
}

func signalDelay(wires []wire) int {
	field := make(map[tupel][]bool)
	paths := make(map[tupel][]int)
	for i, wire := range wires {
		cur := tupel{0, 0}
		p := 0
		for _, move := range wire.ms {
			for j := 0; j < move.dist; j++ {
				switch move.dir {
				case 'R':
					cur.x++
				case 'L':
					cur.x--
				case 'U':
					cur.y++
				case 'D':
					cur.y--
				}
				if _, inUse := field[cur]; !inUse {
					field[cur] = make([]bool, len(wires))
					paths[cur] = make([]int, len(wires))
				}
				field[cur][i] = true
				p++
				paths[cur][i] = p
			}
		}
	}
	distances := []int{}
	for k, v := range field {
		count := 0
		for _, entry := range v {
			if entry {
				count++
			}
		}
		if count > 1 {
			// TODO solve puzzle for more than two wires
			distances = append(distances, paths[k][0]+paths[k][1])
		}
	}
	return min(distances)
}

func min(m []int) int {
	min := 1<<(64-1) - 1
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func readInput(filename string) (wires []wire, err error) {
	wires = make([]wire, 0)
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
			fmt.Sscanf(cur[1:], "%d", &dist)
			ms = append(ms, move{rune(cur[0]), dist})
		}
		wires = append(wires, wire{ms})
	}
	return wires, nil
}
