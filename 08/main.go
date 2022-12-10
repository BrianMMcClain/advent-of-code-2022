package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(path string) []string {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	ret := []string{}
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func parseInput(input []string) [][]int {
	ret := [][]int{}

	for _, line := range input {
		row := []int{}
		for _, c := range line {
			height, _ := strconv.Atoi(string(c))
			row = append(row, height)
		}
		ret = append(ret, row)
	}

	return ret
}

func visibleWest(forest [][]int, x int, y int) (bool, int) {
	i := x - 1
	count := 0
	height := forest[x][y]
	for i >= 0 {
		count++
		if forest[i][y] >= height {
			return false, count
		}
		i--
	}

	return true, count
}

func visibleEast(forest [][]int, x int, y int) (bool, int) {
	i := x + 1
	count := 0
	height := forest[x][y]
	for i < len(forest) {
		count++
		if forest[i][y] >= height {
			return false, count
		}
		i++
	}

	return true, count
}

func visibleNorth(forest [][]int, x int, y int) (bool, int) {
	i := y - 1
	count := 0
	height := forest[x][y]
	for i >= 0 {
		count++
		if forest[x][i] >= height {
			return false, count
		}
		i--
	}

	return true, count
}

func visibleSouth(forest [][]int, x int, y int) (bool, int) {
	i := y + 1
	count := 0
	height := forest[x][y]
	for i < len(forest[x]) {
		count++
		if forest[x][i] >= height {
			return false, count
		}
		i++
	}

	return true, count
}

func part1(forest [][]int) int {
	visible := 0

	for x := 0; x < len(forest); x++ {
		for y := 0; y < len(forest[x]); y++ {
			vw, _ := visibleWest(forest, x, y)
			ve, _ := visibleEast(forest, x, y)
			vn, _ := visibleNorth(forest, x, y)
			vs, _ := visibleSouth(forest, x, y)
			if vw || ve || vn || vs {
				visible++
			}
		}
	}

	return visible
}

func part2(forest [][]int) int {
	maxView := 0
	for x := 0; x < len(forest); x++ {
		for y := 0; y < len(forest[x]); y++ {
			_, vw := visibleWest(forest, x, y)
			_, ve := visibleEast(forest, x, y)
			_, vn := visibleNorth(forest, x, y)
			_, vs := visibleSouth(forest, x, y)
			view := vw * ve * vn * vs
			if view > maxView {
				maxView = view
			}
		}
	}

	return maxView
}

func main() {
	input := readInput("./input.txt")
	forest := parseInput(input)
	fmt.Printf("Part 1: %d\n", part1(forest))
	fmt.Printf("Part 2: %d\n", part2(forest))
}
