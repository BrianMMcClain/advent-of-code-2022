package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filePath string) []string {
	lines := []string{}
	inputFile, _ := os.Open(filePath)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func part1(lines []string, scoreMap map[string]int, beatsMap map[string]string, drawMap map[string]string) int {
	score := 0
	for _, line := range lines {
		round := strings.Split(line, " ")

		// Determine who won
		if drawMap[round[0]] == round[1] {
			score += scoreMap["DRAW"]
		} else if beatsMap[round[0]] == round[1] {
			score += scoreMap["WIN"]
		} else {
			score += scoreMap["LOSE"]
		}

		// Add the score for the shape selected
		score += scoreMap[round[1]]
	}

	return score
}

func part2(lines []string, scoreMap map[string]int, beatsMap map[string]string, drawMap map[string]string) int {

	// Invert the beatsMap to determine how to lose
	losesMap := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	throwMap := map[string]string{
		"X": "LOSE",
		"Y": "DRAW",
		"Z": "WIN",
	}

	score := 0
	for _, line := range lines {
		round := strings.Split(line, " ")

		// Determine if we win, lose, or draw
		roundResult := throwMap[round[1]]
		score += scoreMap[roundResult]

		// Determine which shape we played
		if roundResult == "WIN" {
			score += scoreMap[beatsMap[round[0]]]
		} else if roundResult == "LOSE" {
			score += scoreMap[losesMap[round[0]]]
		} else {
			score += scoreMap[drawMap[round[0]]]
		}
	}

	return score
}

func main() {
	scoreMap := map[string]int{
		"X":    1,
		"Y":    2,
		"Z":    3,
		"WIN":  6,
		"LOSE": 0,
		"DRAW": 3,
	}
	beatsMap := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	drawMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	lines := readInput("./input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines, scoreMap, beatsMap, drawMap))
	fmt.Printf("Part 2: %d\n", part2(lines, scoreMap, beatsMap, drawMap))
}
