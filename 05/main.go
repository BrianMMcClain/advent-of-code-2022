package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Move struct {
	Count int
	From  int
	To    int
}

type Cargo struct {
	Stacks []*list.List
	Moves  []Move
}

func parseMove(sMove string) Move {
	re, _ := regexp.Compile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	m := re.FindStringSubmatch(sMove)

	count, _ := strconv.Atoi(m[1])
	from, _ := strconv.Atoi(m[2])
	to, _ := strconv.Atoi(m[3])

	move := Move{count, from, to}
	return move
}

func parseContainerRow(row string, stackCount int) []string {
	// This can be done easier with a regex but it's getting late
	var ret []string
	index := 0
	for index*4 < stackCount*4 {
		if row[(index*4)+1:(index*4)+2] == " " {
			ret = append(ret, "")
		} else {
			ret = append(ret, row[(index*4)+1:(index*4)+2])
		}
		index++
	}

	return ret
}

func readInput(path string) Cargo {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	// Scan the container stack
	stackCount := 0
	stacks := []*list.List{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		} else if stackCount == 0 {
			// Calculate the number of stacks and initialize the array
			// if we haven't already
			stackCount = (len(scanner.Text()) + 1) / 4
			for i := 0; i < stackCount; i++ {
				stacks = append(stacks, list.New())
			}
		}

		// Get back an array representing a horizontal slice of the stacks
		// If the value is empty, don't add to the stack
		stackRow := parseContainerRow(scanner.Text(), stackCount)
		for i := 0; i < stackCount; i++ {
			if stackRow[i] != "" {
				stacks[i].PushBack(stackRow[i])
			}
		}
	}

	// Ugly way to remove the stack numbering in the input
	for i := 0; i < stackCount; i++ {
		stacks[i].Remove(stacks[i].Back())
	}

	// Parse the moves
	moves := []Move{}
	for scanner.Scan() {
		moves = append(moves, parseMove(scanner.Text()))
	}

	cargo := Cargo{stacks, moves}
	return cargo
}

func topCratesToString(cargo Cargo) string {
	ret := ""
	for i := 0; i < len(cargo.Stacks); i++ {
		v, _ := cargo.Stacks[i].Front().Value.(string)
		ret += v
	}
	return ret
}

func part1(cargo Cargo) string {
	for _, m := range cargo.Moves {
		for i := 0; i < m.Count; i++ {
			e := cargo.Stacks[m.From-1].Front()
			cargo.Stacks[m.From-1].Remove(e)
			cargo.Stacks[m.To-1].PushFront(e.Value)
		}
	}

	return topCratesToString(cargo)
}

func part2(cargo Cargo) string {
	for _, m := range cargo.Moves {
		tmpStack := list.New()
		for i := 0; i < m.Count; i++ {
			e := cargo.Stacks[m.From-1].Front()
			cargo.Stacks[m.From-1].Remove(e)
			tmpStack.PushFront(e.Value)
		}
		for i := 0; i < m.Count; i++ {
			e := tmpStack.Front()
			tmpStack.Remove(e)
			cargo.Stacks[m.To-1].PushFront(e.Value)
		}
	}

	return topCratesToString(cargo)
}

func main() {
	cargo := readInput("./input.txt")
	fmt.Printf("Part 1: %s\n", part1(cargo))

	cargo = readInput("./input.txt")
	fmt.Printf("Part 2: %s\n", part2(cargo))
}
