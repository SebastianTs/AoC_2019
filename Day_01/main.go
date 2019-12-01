package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	masses, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	sumOne, sumTwo := 0, 0
	for _, mass := range masses {
		sumOne += calcFuelRequired(mass)
		sumTwo += calcFuelRequiredPartTwo(mass)
	}
	fmt.Printf("The solution for Part One is %d\n", sumOne)
	fmt.Printf("The solution for Part Two is %d\n", sumTwo)
}

func calcFuelRequired(mass int) int {
	mass /= 3
	mass -= 2
	return mass
}

func calcFuelRequiredPartTwo(mass int) (sum int) {
	for mass > 0 {
		mass = calcFuelRequired(mass)
		if mass > 0 {
			sum += mass
		}
	}
	return sum
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
		mass, err := strconv.Atoi(line)
		if err != nil {
			return ns, nil
		}
		ns = append(ns, mass)
	}
	return ns, nil
}
