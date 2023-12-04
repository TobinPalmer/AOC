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
	maxes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	total := 0

	for i := 0; i < len(strings.Split(input, "\n")); i++ {
		localTotal := 1

		split := strings.Split(input, "\n")[i]
		line := split[strings.Index(split, ":")+1:]
		line = strings.NewReplacer(",", "", ";", "").Replace(line)
		splitLine := strings.Split(strings.ReplaceAll(strings.Trim(line, " "), ";", ""), " ")

		for i := 1; i < len(splitLine); i += 2 {
			v, _ := strconv.Atoi(splitLine[i-1])
			maxes[splitLine[i]] = max(v, maxes[splitLine[i]])
		}

		for _, v := range maxes {
			localTotal *= v
		}
		total += localTotal

		maxes["red"] = 0
		maxes["green"] = 0
		maxes["blue"] = 0
		localTotal = 1
	}

	return total
}
