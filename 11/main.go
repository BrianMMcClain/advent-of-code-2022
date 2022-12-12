package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	ADD      = 1
	MULTIPLY = 2
	SQUARE   = 3
)

type Monkey struct {
	Items           []int
	Operation       int
	OperationValue  int
	TestValue       int
	TrueMonkey      int
	FalseMonkey     int
	InspectionCount int
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

func parseInput(input []string) []Monkey {
	monkeys := []Monkey{}

	monkeyIndex := -1
	for _, line := range input {
		sLine := strings.Split(strings.TrimSpace(line), " ")
		if sLine[0] == "Monkey" {
			// New Monkey
			monkey := Monkey{}
			monkeys = append(monkeys, monkey)
			monkeyIndex++
		} else if sLine[0] == "Starting" {
			// Starting items
			for _, item := range sLine[2:] {
				item = strings.ReplaceAll(item, ",", "")
				iItem, _ := strconv.Atoi(item)
				monkeys[monkeyIndex].Items = append(monkeys[monkeyIndex].Items, iItem)
			}
		} else if sLine[0] == "Operation:" {
			// Test operation
			if sLine[len(sLine)-2] == "+" {
				monkeys[monkeyIndex].Operation = ADD
			} else if sLine[len(sLine)-2] == "*" {
				monkeys[monkeyIndex].Operation = MULTIPLY
			}

			if sLine[len(sLine)-1] == "old" {
				monkeys[monkeyIndex].Operation = SQUARE
			} else {
				opValue, _ := strconv.Atoi(sLine[len(sLine)-1])
				monkeys[monkeyIndex].OperationValue = opValue
			}
		} else if sLine[0] == "Test:" {
			// Test
			testValue, _ := strconv.Atoi(sLine[len(sLine)-1])
			monkeys[monkeyIndex].TestValue = testValue
		} else if sLine[0] != "" && sLine[1] == "true:" {
			throwTo, _ := strconv.Atoi(sLine[len(sLine)-1])
			monkeys[monkeyIndex].TrueMonkey = throwTo
		} else if sLine[0] != "" && sLine[1] == "false:" {
			throwTo, _ := strconv.Atoi(sLine[len(sLine)-1])
			monkeys[monkeyIndex].FalseMonkey = throwTo
		}
	}

	return monkeys
}

func runInspections(monkeys []Monkey, rounds int, worryDiv int, modWorry bool) []Monkey {
	for i := 0; i < rounds; i++ {
		for monkeyIndex := range monkeys {
			for _, item := range monkeys[monkeyIndex].Items {
				worry := item
				if len(monkeys[monkeyIndex].Items) == 1 {
					monkeys[monkeyIndex].Items = []int{}
				} else {
					monkeys[monkeyIndex].Items = monkeys[monkeyIndex].Items[1:]
				}

				// Inspect
				monkeys[monkeyIndex].InspectionCount++
				switch monkeys[monkeyIndex].Operation {
				case ADD:
					if modWorry {
						worry = (worry + monkeys[monkeyIndex].OperationValue) % worryDiv
					} else {
						worry = worry + monkeys[monkeyIndex].OperationValue
					}
				case MULTIPLY:
					if modWorry {
						worry = (worry % worryDiv) * (monkeys[monkeyIndex].OperationValue % worryDiv)
					} else {
						worry = worry * monkeys[monkeyIndex].OperationValue
					}
				case SQUARE:
					if modWorry {
						worry = (worry % worryDiv) * (worry % worryDiv)
					} else {
						worry *= worry
					}
				}

				if !modWorry {
					worry = int(math.Floor(float64(worry) / float64(worryDiv)))
				}

				// Test
				if worry%monkeys[monkeyIndex].TestValue == 0 {
					// True
					monkeys[monkeys[monkeyIndex].TrueMonkey].Items = append(monkeys[monkeys[monkeyIndex].TrueMonkey].Items, worry)
				} else {
					// false
					monkeys[monkeys[monkeyIndex].FalseMonkey].Items = append(monkeys[monkeys[monkeyIndex].FalseMonkey].Items, worry)
				}
			}
		}
	}

	return monkeys
}

func calcMonkeyBusiness(monkeys []Monkey) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].InspectionCount > monkeys[j].InspectionCount
	})
	return monkeys[0].InspectionCount * monkeys[1].InspectionCount
}

func main() {
	input := readInput("./input.txt")
	monkeys := parseInput(input)

	p1Monkeys := make([]Monkey, len(monkeys))
	copy(p1Monkeys, monkeys)
	runInspections(p1Monkeys, 20, 3, false)
	fmt.Printf("Part 1: %d\n", calcMonkeyBusiness(p1Monkeys))

	p2Monkeys := make([]Monkey, len(monkeys))
	copy(p2Monkeys, monkeys)
	worryDiv := 1
	for _, m := range p2Monkeys {
		worryDiv *= m.TestValue
	}
	runInspections(p2Monkeys, 10000, worryDiv, true)
	fmt.Printf("Part 2: %d\n", calcMonkeyBusiness(p2Monkeys))
}
