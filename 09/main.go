package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func calculateTail(hX, hY, tX, tY int) (int, int) {
	diffX := hX - tX
	diffY := hY - tY
	movedX, movedY := false, false

	if diffX == 2 {
		tX++
		movedX = true
	} else if diffX == -2 {
		tX--
		movedX = true
	}

	if diffY == 2 {
		tY++
		movedY = true
	} else if diffY == -2 {
		tY--
		movedY = true
	}

	if movedX && (diffY != 0) && !movedY {
		tY += diffY
	} else if movedY && (diffX != 0) && !movedX {
		tX += diffX
	}

	return tX, tY
}

type Position struct {
	X int
	Y int
}

func calculateMovement(input []string, knotCount int) int {
	visited := []map[string]int{}
	knots := []Position{}
	for i := 0; i < knotCount; i++ {
		visited = append(visited, make(map[string]int))
		visited[i]["0,0"] = 1
		pos := Position{0, 0}
		knots = append(knots, pos)
	}

	for _, line := range input {
		sLine := strings.Split(line, " ")
		dir := sLine[0]
		count, _ := strconv.Atoi(sLine[1])

		for count > 0 {
			switch dir {
			case "L":
				knots[0].X--
			case "R":
				knots[0].X++
			case "U":
				knots[0].Y--
			case "D":
				knots[0].Y++
			}

			for i := 1; i < knotCount; i++ {
				tX, tY := calculateTail(knots[i-1].X, knots[i-1].Y, knots[i].X, knots[i].Y)
				knots[i].X = tX
				knots[i].Y = tY
				visited[i][fmt.Sprintf("%d,%d", tX, tY)]++

			}

			count--
		}
	}

	return len(visited[knotCount-1])
}

func main() {
	input := readInput("./input.txt")
	fmt.Printf("Part 1: %d\n", calculateMovement(input, 2))
	fmt.Printf("Part 1: %d\n", calculateMovement(input, 10))
}
