package main

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"os"
	"slices"
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
	split := strings.Split(input, "\n\n")
	seeds := stringArrToIntArr(strings.Split(split[0][7:], " "))

	for i := 1; i < len(split); i++ {
		inputs := spanningStringToIntArr(strings.Split(strings.TrimSpace(split[i][strings.Index(split[i], ":")+1:]), "\n"))
		seeds = generateMutation(inputs, seeds)
	}

	return slices.Min(seeds)
}

func stringArrToIntArr(input []string) []int {
	var output []int

	for i := 0; i < len(input); i++ {
		val, err := strconv.Atoi(input[i])
		if err != nil {
			fmt.Println(err)
			continue
		}

		output = append(output, val)
	}

	return output
}

func spanningStringToIntArr(input []string) [][]int {
	var output [][]int

	for i := 0; i < len(input); i++ {
		innerSplit := strings.Split(input[i], " ")
		var tempArr []int

		for j := 0; j < len(innerSplit); j++ {
			val, err := strconv.Atoi(innerSplit[j])
			if err != nil {
				fmt.Println("error parsing", err)
			}

			tempArr = append(tempArr, val)
		}

		output = append(output, tempArr)
	}

	return output
}

// Returns the new seed values
func generateMutation(key [][]int, seed []int) []int {
	var output []int

outer:
	for i := 0; i < len(seed); i++ {
		for j := 0; j < len(key); j++ {
			if key[j][1] <= seed[i] && seed[i] <= key[j][1]+key[j][2] {
				output = append(output, seed[i]+(key[j][0]-key[j][1]))
				continue outer
			}
		}

		output = append(output, seed[i])
	}

	return output
}

func part2(input string) int {
	split := strings.Split(input, "\n\n")
	seedsArr := stringArrToIntArr(strings.Split(split[0][7:], " "))
	var seeds []int

	for i := 0; i < len(seedsArr)-1; i += 2 {
		for j := 0; j < seedsArr[i+1]; j++ {
			seeds = append(seeds, seedsArr[i]+j)
		}
	}

	for i := 1; i < len(split); i++ {
		inputs := spanningStringToIntArr(strings.Split(strings.TrimSpace(split[i][strings.Index(split[i], ":")+1:]), "\n"))
		seeds = generateMutation(inputs, seeds)
	}

	return slices.Min(seeds)
}
