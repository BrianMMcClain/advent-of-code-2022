package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getList(inputFile string) []int {
	fInput, _ := os.Open(inputFile)
	defer fInput.Close()

	currentCalories := 0
	var elfList []int
	scanner := bufio.NewScanner(fInput)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfList = append(elfList, currentCalories)
			currentCalories = 0
		} else {
			iCalories, _ := strconv.Atoi(line)
			currentCalories += iCalories
		}
	}

	sort.Ints(elfList)
	return elfList
}

func part1(elfList []int) int {
	return elfList[len(elfList)-1]
}

func part2(elfList []int) int {
	elfSlice := elfList[len(elfList)-3 : len(elfList)]

	sum := 0
	for _, calories := range elfSlice {
		sum += calories
	}
	return sum
}

func main() {
	elfList := getList("./input.txt")
	fmt.Printf("Part 1: %d\n", part1(elfList))
	fmt.Printf("Part 2: %d\n", part2(elfList))
}
