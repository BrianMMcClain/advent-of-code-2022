package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Beacon struct {
	x int
	y int
}

type Sensor struct {
	x        int
	y        int
	beacon   Beacon
	distance int
}

type SensorNetwork struct {
	sensors []Sensor
	beacons []Beacon
}

const (
	EMPTY   = 0
	SENSOR  = 1
	BEACON  = 2
	VISITED = 3
)

func readInput(path string) []string {
	ret := []string{}
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func distance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func parseInput(input []string) SensorNetwork {
	net := SensorNetwork{}

	r, _ := regexp.Compile("Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)")

	for _, line := range input {
		parsed := r.FindStringSubmatch(line)

		sensor := Sensor{}
		sensor.x, _ = strconv.Atoi(parsed[1])
		sensor.y, _ = strconv.Atoi(parsed[2])

		beacon := Beacon{}
		beacon.x, _ = strconv.Atoi(parsed[3])
		beacon.y, _ = strconv.Atoi(parsed[4])

		sensor.distance = distance(sensor.x, sensor.y, beacon.x, beacon.y)
		sensor.beacon = beacon

		net.sensors = append(net.sensors, sensor)
		net.beacons = append(net.beacons, beacon)
	}

	return net
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

func part1(net *SensorNetwork, row int) map[int]int {
	filledRow := map[int]int{}
	for _, s := range net.sensors {
		if s.distance >= abs(s.y-row) {
			// Line between beacon and sensor intersects with the row of interest
			fillWidth := s.distance - abs(s.y-row)
			for x := s.x - fillWidth; x <= s.x+fillWidth; x++ {
				filledRow[x] = VISITED
			}
		}

		if s.y == row {
			filledRow[s.x] = SENSOR
		}
		if s.beacon.y == row {
			filledRow[s.beacon.x] = BEACON
		}
	}

	return filledRow
}

func main() {
	input := readInput("./input.txt")
	net := parseInput(input)

	p1row := part1(&net, 2000000)
	count := 0
	for _, v := range p1row {
		if v != BEACON {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}
