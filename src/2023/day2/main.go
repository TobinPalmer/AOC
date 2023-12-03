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
	const codePointToNumMagic uint8 = 48
	total := 0

	for _, line := range strings.Split(input, "\n") {
		var firstNum uint8 = 0
		var lastNum uint8 = 0

		for i := 0; i < len(line); i++ {
			if 48 <= line[i] && line[i] <= 57 {
				firstNum = line[i] - codePointToNumMagic
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if 48 <= line[i] && line[i] <= 57 {
				lastNum = line[i] - codePointToNumMagic
				break
			}
		}

		total += int(firstNum*10 + lastNum)
	}

	return total
}

func part2(input string) int {
	const codePointToNumMagic uint8 = 48
	stringsToNum := map[string]uint8{"nine": 9, "eight": 8, "seven": 7, "six": 6, "five": 5, "four": 4, "three": 3, "two": 2, "one": 1, "zero": 0}
	total := 0

	for _, line := range strings.Split(input, "\n") {
		var firstNum uint8 = 0
		var lastNum uint8 = 0

		res := ""
	o1:
		for i := 0; i < len(line); i++ {
			res += string(line[i])

			for k, v := range stringsToNum {
				if len(res) < len(k) {
					continue
				}

				if res[len(res)-len(k):] == k {
					firstNum = v
					break o1
				}
			}

			if 48 <= line[i] && line[i] <= 57 {
				firstNum = line[i] - codePointToNumMagic
				break o1
			}

		}

		resEnd := ""
	o2:
		for i := len(line) - 1; i >= 0; i-- {
			resEnd = line[i:]

			for k, v := range stringsToNum {
				if len(resEnd) < len(k) {
					continue
				}
				if resEnd[:len(k)] == k {
					lastNum = v
					break o2
				}
			}

			if 48 <= line[i] && line[i] <= 57 {
				lastNum = line[i] - codePointToNumMagic
				break o2
			}
		}
		total += int(firstNum*10 + lastNum)
	}

	return total
}
