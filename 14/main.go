package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type CaveMap struct {
	walls   [][]Point
	maxX    int
	maxY    int
	caveMap [][]int
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

func parseLine(line string) []Point {
	ret := []Point{}
	sLine := strings.Split(line, " -> ")
	for _, p := range sLine {
		sPoint := strings.Split(p, ",")
		x, _ := strconv.Atoi(sPoint[0])
		y, _ := strconv.Atoi(sPoint[1])
		point := Point{x, y}
		ret = append(ret, point)
	}
	return ret
}

const (
	AIR  = 0
	ROCK = 1
	SAND = 2
)

func drawMap(input []string) CaveMap {
	cMap := CaveMap{}
	cMap.maxX = 500
	cMap.maxY = 0

	for _, line := range input {
		cMap.walls = append(cMap.walls, parseLine(line))
	}

	for i := 0; i < len(cMap.walls); i++ {
		for j := 0; j < len(cMap.walls[i]); j++ {
			if cMap.walls[i][j].x > cMap.maxX {
				cMap.maxX = cMap.walls[i][j].x + 1
			}
			if cMap.walls[i][j].y > cMap.maxY {
				cMap.maxY = cMap.walls[i][j].y + 1
			}
		}
	}

	// Initalize the map
	for y := 0; y < cMap.maxY; y++ {
		cMap.caveMap = append(cMap.caveMap, []int{})
		for x := 0; x < cMap.maxX; x++ {
			cMap.caveMap[y] = append(cMap.caveMap[y], AIR)
		}
	}

	// Draw the walls
	for _, line := range cMap.walls {
		x, y := line[0].x, line[0].y
		cMap.caveMap[y][x] = ROCK
		for _, point := range line[1:] {
			xStep, yStep := 0, 0
			if x-point.x < 0 {
				xStep = 1
			} else if x-point.x > 0 {
				xStep = -1
			}

			if y-point.y < 0 {
				yStep = 1
			} else if y-point.y > 0 {
				yStep = -1
			}

			stepCount := int(math.Abs(float64((x - point.x) + (y - point.y))))
			for i := 0; i < stepCount; i++ {
				x += xStep
				y += yStep
				cMap.caveMap[y][x] = ROCK
			}
		}
	}

	return cMap
}

func simulate(cMap CaveMap, floor bool) []Point {
	offMap := false
	newSand := true

	sands := []Point{}

	sandX, sandY := 500, 0
	for !offMap {
		if newSand {
			sandX, sandY = 500, 0
			newSand = false
		} else {
			if sandY+1 >= cMap.maxY {
				// Into the abyss
				return sands
			} else if sandX+1 >= cMap.maxX {
				// Expand one to the right
				for y := 0; y < cMap.maxY; y++ {
					cMap.caveMap[y] = append(cMap.caveMap[y], AIR)
				}
				if floor {
					cMap.caveMap[cMap.maxY-1][cMap.maxX] = ROCK
				}
				cMap.maxX++
			} else if cMap.caveMap[sandY+1][sandX] == AIR {
				// Nothing below
				sandY++
			} else if cMap.caveMap[sandY+1][sandX-1] == AIR {
				// Roll Left
				sandY++
				sandX--
			} else if cMap.caveMap[sandY+1][sandX+1] == AIR {
				// Roll Right
				sandY++
				sandX++
			} else {
				newSand = true
				sands = append(sands, Point{sandX, sandY})
				cMap.caveMap[sandY][sandX] = SAND
				if sandX == 500 && sandY == 0 {
					return sands
				}
			}
		}
	}

	return sands
}

func printMap(cMap CaveMap) {
	viz := []rune{'.', '#', 'o'}
	for y := 0; y < cMap.maxY; y++ {
		for x := 486; x < cMap.maxX; x++ {
			fmt.Print(string(viz[cMap.caveMap[y][x]]))
		}
		fmt.Println()
	}
}

func main() {
	input := readInput("./input.txt")
	p1CaveMap := drawMap(input)
	p1Sands := simulate(p1CaveMap, false)
	fmt.Printf("Part 1: %d\n", len(p1Sands))

	p2CaveMap := drawMap(input)
	p2CaveMap.caveMap = append(p2CaveMap.caveMap, []int{})
	p2CaveMap.caveMap = append(p2CaveMap.caveMap, []int{})
	for i := 0; i < len(p2CaveMap.caveMap[0]); i++ {
		p2CaveMap.caveMap[len(p2CaveMap.caveMap)-2] = append(p2CaveMap.caveMap[len(p2CaveMap.caveMap)-2], AIR)
		p2CaveMap.caveMap[len(p2CaveMap.caveMap)-1] = append(p2CaveMap.caveMap[len(p2CaveMap.caveMap)-1], ROCK)
	}
	p2CaveMap.maxY += 2

	p2Sands := simulate(p2CaveMap, true)
	fmt.Printf("Part 2: %d\n", len(p2Sands))
}
