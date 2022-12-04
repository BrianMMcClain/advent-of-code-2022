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
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseLine(line string) (int, int, int, int) {
	splitLine := strings.Split(line, ",")
	a := strings.Split(splitLine[0], "-")
	al, _ := strconv.Atoi(a[0])
	ar, _ := strconv.Atoi(a[1])
	b := strings.Split(splitLine[1], "-")
	bl, _ := strconv.Atoi(b[0])
	br, _ := strconv.Atoi(b[1])
	return al, ar, bl, br
}

func part1(lines []string) int {
	totalContains := 0

	for _, line := range lines {
		al, ar, bl, br := parseLine(line)
		if (al <= bl && ar >= br) || (bl <= al && br >= ar) {
			totalContains++
		}
	}
	return totalContains
}

func overlaps(al, ar, bl, br int) bool {
	if al == bl || al == br || ar == bl || al == br {
		return true
	}

	if al < bl && ar > bl {
		return true
	}

	if al > bl && al < br {
		return true
	}

	return false
}

func part2(lines []string) int {
	totalOverlaps := 0

	for _, line := range lines {
		al, ar, bl, br := parseLine(line)
		if overlaps(al, ar, bl, br) || overlaps(bl, br, al, ar) {
			totalOverlaps++
		}
	}

	return totalOverlaps
}

func main() {
	lines := readInput("./input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
