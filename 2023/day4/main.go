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
	total := 0

	for i := 0; i < len(split); i++ {
		localTotal := 0

		line := strings.Trim(split[i][strings.Index(split[i], ":")+1:], " ")
		winningNums := strings.Split(strings.Trim(strings.Split(line, "|")[0], " "), " ")
		inputNums := strings.Split(strings.Trim(strings.Split(line, "|")[1], " "), " ")
		winningNumsMap := make(map[uint]struct{})
		inputNumsMap := make(map[uint]struct{})

		for _, val := range winningNums {
			v, err := strconv.Atoi(strings.TrimSpace(val))
			if err == nil {
				winningNumsMap[uint(v)] = struct{}{}
			}
		}

		for _, val := range inputNums {
			v, err := strconv.Atoi(strings.TrimSpace(val))
			if err == nil {
				inputNumsMap[uint(v)] = struct{}{}
			}
		}

		fmt.Println(winningNumsMap)

		for k := range inputNumsMap {
			_, ok := winningNumsMap[k]
			if !ok {
				continue
			}

			if localTotal == 0 {
				localTotal++
			} else {
				localTotal *= 2
			}
		}

		total += localTotal
		localTotal = 0
	}

	return total
}

func part2(input string) int {
	fmt.Println(input)
	return 0
}
