package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	digits = 5
)

func main() {
	ns, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	original := make([]int, len(ns))
	copy(original, ns)
	ns[1] = 1

	processIntcode(ns)
	fmt.Printf("Part1: Solution:\t%d\n", ns[0])
}

func processIntcode(ns []int) {
	ip := 0
	a, b, c := 0, 0, 0
	opCodeLength := 4
	for ns[ip] != 99 {
		switch ns[ip] {
		case 1:
			a = ns[ip+1]
			b = ns[ip+2]
			c = ns[ip+3]
			ns[c] = ns[a] + ns[b]
			opCodeLength = 4
		case 2:
			a = ns[ip+1]
			b = ns[ip+2]
			c = ns[ip+3]
			ns[c] = ns[a] * ns[b]
			opCodeLength = 4
		case 3:
			//TODO
			opCodeLength = 2
		case 4:
			//TODO
			opCodeLength = 2
		default:
			//TODO
		}
		ip += opCodeLength
	}
}

func readInput(filename string) (ns []int, err error) {
	ns = make([]int, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ns, err
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
		for scntmp.Scan() {
			n, err := strconv.Atoi(scntmp.Text())
			if err != nil {
				return ns, err
			}
			ns = append(ns, n)
		}
	}
	return ns, nil
}

func intToArray(number int) (out [digits]int) {
	for i := 0; i < digits; i++ {
		out[digits-i-1] = number % 10
		number /= 10
	}
	return out
}
