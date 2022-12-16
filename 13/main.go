package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

func compareFloats(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func compare(a, b any) int {
	aArr, aIsArr := a.([]any)
	bArr, bIsArr := b.([]any)

	if !aIsArr && !bIsArr {
		return compareFloats(a.(float64), b.(float64))
	} else if !aIsArr {
		aArr = []any{a}
	} else if !bIsArr {
		bArr = []any{b}
	}

	for i := 0; i < len(aArr) && i < len(bArr); i++ {
		c := compare(aArr[i], bArr[i])
		if c != 0 {
			return c
		}
	}

	return len(aArr) - len(bArr)
}

func main() {
	input := readInput("./input.txt")

	correctIndexSum := 0
	packets := []any{}

	for i := 0; i < (len(input)+1)/3; i++ {
		as := input[i*3]
		bs := input[(i*3)+1]

		var a any
		json.Unmarshal([]byte(as), &a)
		var b any
		json.Unmarshal([]byte(bs), &b)
		packets = append(packets, a, b)
		c := compare(a, b)
		if c <= 0 {
			correctIndexSum += i + 1
		}
	}

	var div1 any
	json.Unmarshal([]byte("[[2]]"), &div1)
	var div2 any
	json.Unmarshal([]byte("[[6]]"), &div2)

	packets = append(packets, div1, div2)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) <= 0
	})

	div1i, div2i := 0, 0
	for i := 0; i < len(packets); i++ {
		if compare(packets[i], div1) == 0 {
			div1i = i + 1
		} else if compare(packets[i], div2) == 0 {
			div2i = i + 1
		}
	}

	fmt.Printf("Part 1: %d\n", correctIndexSum)
	fmt.Printf("Part 2: %d\n", div1i*div2i)
}
