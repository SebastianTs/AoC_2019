package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	opCodeLength = 4
)

func main() {
	ns, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	original := make([]int, len(ns))
	copy(original, ns)
	//Part 1
	ns[1] = 12
	ns[2] = 2
	processIntcode(ns)
	fmt.Printf("Part1: Solution:\t%d\n", ns[0])
	//Part 2
	noun, verb, err := findOutput(19690720, original)
	if err != nil {
		panic(err)
	}
	fmt.Printf("noun %d, verb %d\n", noun, verb)
	fmt.Printf("Part2: Solution:\t%d\n", 100*noun+verb)
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
	for ns[ip] != 99 {
		switch ns[ip] {
		/*Opcode 1 adds together numbers read from two positions and stores the result in a third position.
		The three integers immediately after the opcode tell you these three positions -
		the first two indicate the positions from which you should read the input values,
		and the third indicates the position at which the output should be stored.*/
		case 1:
			a = ns[ip+1]
			b = ns[ip+2]
			c = ns[ip+3]
			ns[c] = ns[a] + ns[b]
		/*Opcode 2 works exactly like opcode 1,
		except it multiplies the two inputs instead of adding them.
		Again, the three integers after the opcode indicate where the inputs and outputs are,
		not their values.*/
		case 2:
			a = ns[ip+1]
			b = ns[ip+2]
			c = ns[ip+3]
			ns[c] = ns[a] * ns[b]
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
