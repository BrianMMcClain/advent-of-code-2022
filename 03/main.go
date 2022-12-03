package main

import (
	"bufio"
	"fmt"
	"os"
)

func intersection(l1, l2 string) []rune {
	runeMap := map[rune]int{}
	common := []rune{}
	for _, r := range l1 {
		runeMap[r]++
	}
	for _, r := range l2 {
		if runeMap[r] > 0 {
			common = append(common, r)
		}
	}

	return common
}

func readInput(filePath string) []string {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	} else if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}

	return -1
}

func part1(lines []string) int {
	prio := 0
	for _, line := range lines {
		h1 := line[0 : len(line)/2]
		h2 := line[len(line)/2:]
		i := intersection(h1, h2)
		prio += priority(i[0])
	}

	return prio
}

func part2(lines []string) int {
	prio := 0
	index := 0

	for index+2 <= len(lines) {
		r1 := intersection(lines[index], lines[index+1])
		r2 := intersection(lines[index+2], string(r1))
		prio += priority(r2[0])
		index += 3
	}

	return prio

}

func main() {
	lines := readInput("./input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
