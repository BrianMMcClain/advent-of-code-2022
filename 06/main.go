package main

import (
	"fmt"
	"os"
)

func uniq(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}

	return true
}

func part1(input string, markerLen int) int {
	for i := 0; i < len(input)-markerLen+1; i++ {
		if uniq(input[i : i+markerLen]) {
			return i + markerLen
		}
	}

	return -1
}

func main() {
	bin, _ := os.ReadFile("./input.txt")
	input := string(bin)
	fmt.Println(part1(input, 4))
	fmt.Println(part1(input, 14))
}
