package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Location struct {
	X       int
	Y       int
	Height  int
	Visited bool
	IsStart bool
	IsEnd   bool
	Parent  *Location
}

type Land struct {
	Map   [][]Location
	Start *Location
	End   *Location
}

func readInput(path string) []string {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	ret := []string{}

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func buildLand(input []string) Land {
	land := Land{}

	for y, line := range input {
		landLine := []Location{}
		for x, r := range line {
			l := Location{}
			l.Height = -1
			if r == 'S' {
				l.Height = 0
				l.IsStart = true
				l.IsEnd = false
				l.Visited = true
				land.Start = &l
			} else if r == 'E' {
				l.Height = int('z') - int('a') + 2 // Height of z + 1
				l.IsStart = false
				l.IsEnd = true
				land.End = &l
			} else {
				l.IsStart = false
				l.IsEnd = false
				l.Height = int(r) - int('a') + 1
			}
			l.X = x
			l.Y = y

			landLine = append(landLine, l)
		}
		land.Map = append(land.Map, landLine)
	}

	return land
}

func toVisit(land [][]int, visited [][]bool, x int, y int) bool {
	if x < 0 || y < 0 || y > len(land) || x > len(land[y]) {
		return false
	} else if visited[y][x] {
		return false
	}
	return true
}

func visit(land Land) {
}

func dequeue(queue []*Location) (*Location, []*Location) {
	l := queue[0]

	if len(queue) == 1 {
		queue = []*Location{}
	} else {
		queue = queue[1:]
	}

	return l, queue
}

func getConnected(land Land, curLoc *Location) []*Location {
	ret := []*Location{}
	if curLoc.Y > 0 {
		// Up
		if !land.Map[curLoc.Y-1][curLoc.X].Visited && land.Map[curLoc.Y-1][curLoc.X].Height-curLoc.Height <= 1 {
			ret = append(ret, &land.Map[curLoc.Y-1][curLoc.X])
		}
	}
	if curLoc.Y < len(land.Map)-1 {
		// Down
		if !land.Map[curLoc.Y+1][curLoc.X].Visited && land.Map[curLoc.Y+1][curLoc.X].Height-curLoc.Height <= 1 {
			ret = append(ret, &land.Map[curLoc.Y+1][curLoc.X])
		}
	}
	if curLoc.X > 0 {
		// Left
		if !land.Map[curLoc.Y][curLoc.X-1].Visited && land.Map[curLoc.Y][curLoc.X-1].Height-curLoc.Height <= 1 {
			ret = append(ret, &land.Map[curLoc.Y][curLoc.X-1])
		}
	}
	if curLoc.X < len(land.Map[curLoc.Y])-1 {
		// Right
		if !land.Map[curLoc.Y][curLoc.X+1].Visited && land.Map[curLoc.Y][curLoc.X+1].Height-curLoc.Height <= 1 {
			ret = append(ret, &land.Map[curLoc.Y][curLoc.X+1])
		}
	}
	return ret
}

func countSteps(l *Location) int {
	i := l
	steps := 0
	for i.Parent != nil {
		steps++
		i = i.Parent
	}
	return steps
}

func findPath(land Land, start *Location, endElevation int) int {
	queue := []*Location{}
	start.Visited = true
	queue = append(queue, start)
	var l *Location
	for len(queue) > 0 {
		l, queue = dequeue(queue)
		if l.Height == endElevation {
			return countSteps(l)
		}
		for _, next := range getConnected(land, l) {
			next.Visited = true
			next.Parent = l
			queue = append(queue, next)
		}
	}

	return math.MaxInt32
}

func resetLand(land *Land) {
	for y := 0; y < len(land.Map); y++ {
		for x := 0; x < len(land.Map[y]); x++ {
			land.Map[y][x].Visited = false
			land.Map[y][x].Parent = nil
		}
	}
}

func main() {
	input := readInput("./input.txt")
	land := buildLand(input)

	fmt.Printf("Part 1: %d \n", findPath(land, land.Start, 27))

	startingSpots := []*Location{}
	for y := 0; y < len(land.Map); y++ {
		for x := 0; x < len(land.Map[y]); x++ {
			if land.Map[y][x].Height == 1 {
				startingSpots = append(startingSpots, &land.Map[y][x])
			}
		}
	}

	minDistance := math.MaxInt32
	for _, s := range startingSpots {
		resetLand(&land)
		distance := findPath(land, s, 27)
		if distance < minDistance {
			minDistance = distance
		}
	}

	fmt.Printf("Part 2: %d\n", minDistance)
}
