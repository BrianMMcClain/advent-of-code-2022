package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
}

func readInput(path string) []string {
	ret := []string{}
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func getStrengthAndChecks(cycle int, x int, curStrength int, cycleChecks []int) (int, []int) {
	if len(cycleChecks) > 0 && cycle == cycleChecks[0] {
		strength := x * cycle
		cycleChecks = cycleChecks[1:]

		return curStrength + strength, cycleChecks
	}

	return curStrength, cycleChecks
}

func drawCRT(cycle int, x int, crt [][]string) [][]string {

	row := int(math.Floor(float64(cycle-1) / 40.0))
	col := cycle - (row * 40) - 1

	if col >= x-1 && col <= x+1 {
		crt[row][col] = "#"
	}
	return crt
}

func part1and2(input []string, cycleChecks []int) (int, [][]string) {

	x := 1
	cycle := 1
	strengthSum := 0

	// Initialize the CRT
	crt := [][]string{}
	for y := 0; y < 6; y++ {
		crt = append(crt, []string{})
		for x := 0; x < 40; x++ {
			crt[y] = append(crt[y], ".")
		}
	}

	for _, line := range input {
		crt = drawCRT(cycle, x, crt)
		strengthSum, cycleChecks = getStrengthAndChecks(cycle, x, strengthSum, cycleChecks)

		if line == "noop" {
			cycle++
		} else {
			sLine := strings.Split(line, " ")
			if sLine[0] == "addx" {
				cycle++
				crt = drawCRT(cycle, x, crt)
				strengthSum, cycleChecks = getStrengthAndChecks(cycle, x, strengthSum, cycleChecks)

				y, _ := strconv.Atoi(sLine[1])
				x += y
				cycle++
			}
		}
	}

	return strengthSum, crt
}

func main() {
	input := readInput("./input.txt")
	sumStrength, crt := part1and2(input, []int{20, 60, 100, 140, 180, 220})
	fmt.Printf("Part 1: %d\n", sumStrength)

	// Print CRT
	fmt.Println("Part 2:")
	for i := 0; i < len(crt); i++ {
		for j := 0; j < len(crt[i]); j++ {
			fmt.Print(crt[i][j])
		}
		fmt.Println()
	}
}
