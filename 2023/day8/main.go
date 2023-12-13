package main

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	part, err := strconv.Atoi(os.Args[1])
	if err != nil {
		part = 1
	}

	if part == 1 {
		ans := part1(input)
		err := util.CopyToClipboard(fmt.Sprintf("%v", ans))
		if err != nil {
			panic(err)
		}

		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		err := util.CopyToClipboard(fmt.Sprintf("%v", ans))
		if err != nil {
			panic(err)
		}

		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	split := strings.Split(input, "\n")
	pattern := split[0]
	mapper := make(map[string][]string)
	currentString := "AAA"
	total := 0

	for i := 2; i < len(split); i++ {
		startNode := split[i][:3]
		mapping := strings.Split(split[i][7:15], ", ")
		mapper[startNode] = mapping
	}

	patternIndex := 0
	for currentString != "ZZZ" {
		switch string(pattern[patternIndex]) {
		case "L":
			currentString = mapper[currentString][0]
		case "R":
			currentString = mapper[currentString][1]
		}
		total++

		if patternIndex >= len(pattern)-1 {
			patternIndex = 0
		} else {
			patternIndex++
		}
	}

	fmt.Println(mapper)

	return total
}

func part2(input string) int {
	split := strings.Split(input, "\n")

	for i := 0; i < len(split); i++ {
	}

	return 0
}
