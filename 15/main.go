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
	beacon   *Beacon
	distance int
}

type SensorNetwork struct {
	sensors []Sensor
	beacons []Beacon
}

type FillBlock struct {
	start  int
	stop   int
	merged bool
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
		sensor.beacon = &beacon

		net.sensors = append(net.sensors, sensor)

		// Don't double-add beacons
		addBeacon := true
		for _, b := range net.beacons {
			if beacon.x == b.x && beacon.y == b.y {
				addBeacon = false
			}
		}
		if addBeacon {
			net.beacons = append(net.beacons, beacon)
		}
	}

	return net
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

func part1(net *SensorNetwork, row int) []*FillBlock {
	blocks := []*FillBlock{}
	for _, s := range net.sensors {
		if s.distance >= abs(s.y-row) {
			// Line between beacon and sensor intersects with the row of interest
			fillWidth := s.distance - abs(s.y-row)
			b := FillBlock{s.x - fillWidth, s.x + fillWidth, false}
			blocks = append(blocks, &b)
		}

		// Include the sensor if it's on the row as it may not be visited
		if s.y == row {
			blocks = append(blocks, &FillBlock{s.x, s.x, false})
		}
	}

	retBlocks := mergeBlocks(blocks)

	return retBlocks
}

func part2(net *SensorNetwork) int {
	for y := 0; y < 4000000; y++ {
		row := part1(net, y)
		if len(row) > 1 {
			x, y := 0, y
			if row[0].start < row[1].start {
				// In order
				x = row[0].stop + 1
			} else {
				x = row[1].stop + 1
			}
			return (x * 4000000) + y
		}
	}

	return -1
}

func mergeBlocks(blocks []*FillBlock) []*FillBlock {
	mergeBlock := blocks[0]
	blocks[0].merged = true

	anyMerge := false
	mergePerformed := true
	for mergePerformed {
		mergePerformed = false
		for bi, b := range blocks[1:] {
			lMerge, rMerge, fMerge := false, false, false
			if !b.merged && b.start <= mergeBlock.start && b.stop >= mergeBlock.start-1 {
				// Overlap to the left
				mergeBlock.start = b.start
				lMerge = true
				mergePerformed = true
				anyMerge = true
			}
			if !b.merged && b.start <= mergeBlock.stop+1 && b.stop >= mergeBlock.stop {
				// Overlap to the right
				mergeBlock.stop = b.stop
				rMerge = true
				mergePerformed = true
				anyMerge = true
			}
			if !b.merged && b.start >= mergeBlock.start && b.stop <= mergeBlock.stop {
				fMerge = true
				mergePerformed = true
				anyMerge = true
			}

			if lMerge || rMerge || fMerge {
				blocks[bi+1].merged = true
			}
		}
	}

	unmergedBlocks := []*FillBlock{}
	for _, b := range blocks {
		if !b.merged {
			unmergedBlocks = append(unmergedBlocks, b)
		}
	}

	if !anyMerge || len(unmergedBlocks) <= 1 {
		return append(unmergedBlocks, mergeBlock)
	} else {
		return append(mergeBlocks(unmergedBlocks), mergeBlock)
	}
}

func main() {
	input := readInput("./input.txt")
	net := parseInput(input)

	row := 2000000
	p1Blocks := part1(&net, row)

	count := 0
	for _, b := range p1Blocks {
		count += b.stop - b.start + 1
		for _, beacon := range net.beacons {
			if beacon.y == row && b.start <= beacon.x && b.stop >= beacon.x {
				count--
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)

	fmt.Printf("Part 2: %d\n", part2(&net))
}
