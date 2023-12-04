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
	total := 0
	maxes := map[string]uint8{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	//arr := make([]uint8, 3)
outer:
	for i := 0; i < len(strings.Split(input, "\n")); i++ {
		split := strings.Split(input, "\n")[i]
		line := split[strings.Index(split, ":")+1:]

		segments := strings.Split(line, ";")
		for i := 0; i < len(segments); i++ {
			subSegments := strings.Split(segments[i], ",")

			for j := 0; j < len(subSegments); j++ {
				newSplit := strings.Split(strings.Trim(subSegments[j], " "), " ")
				val, _ := strconv.Atoi(newSplit[0])

				if maxes[newSplit[1]] < uint8(val) {
					continue outer
				}
			}
		}

		total += i + 1
	}

	return total
}

func part2(input string) int {
	fmt.Println(input)
	return 0
}
