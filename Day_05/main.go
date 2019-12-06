package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ns, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	original := make([]int, len(ns))
	copy(original, ns)
	//Part 1
	ns[1] = 1
	processIntcode(ns)
	fmt.Printf("Part1: Solution:\t%d\n", ns[0])
}

func findOutput(output int, ns []int) (noun, verb int, err error) {
	original := make([]int, len(ns))
	copy(original, ns)
	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			copy(ns, original)
			ns[1] = noun
			ns[2] = verb
			processIntcode(ns)
			if ns[0] == output {
				return noun, verb, nil
			}
		}
	}
	return noun, verb, errors.New("noun and verb to given output not found")
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

		case 2:
			a = ns[ip+1]
			b = ns[ip+2]
			c = ns[ip+3]
			ns[c] = ns[a] * ns[b]
		
		case 3:
			a = ns[ip+1]
			ns[a] = a
			opCodeLength = 2

		case 4:
			a = ns[ip+1]
			fmt.Println(ns[a])
			opCodeLength = 2
		default:
			cur := ns[ip]
			in := []int{}
			for i:=1; cur > 0;i++{
				in = append(foo,cur%(10*i))
				cur /= 10
			}
			switch in[len(in)-1]{
				case 1:
					a = ns[ip+1]
					b = ns[ip+2]
					c = ns[ip+3]
					ns[c] = ns[a] + ns[b]
				case 2:
				case 3:
				case 4:
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
