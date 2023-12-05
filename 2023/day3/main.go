package main

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"os"
	"regexp"
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
	//sum := 0

	for line := 0; line < len(split); line++ {
		for char := 0; char < len(split[line]); char++ {
			if isNumber(string(split[line][char])) {
				fmt.Println(getNumberLength(line, char, split))
			}
		}
	}

	return 0
}

func getNumberLength(startLine, startIdx int, input []string) (number int, length int) {
	for isNumber(string(input[startLine][startIdx])) {

	}

	return
}

func isSpecial(ch string) bool {
	return !regexp.MustCompile(`\d`).MatchString(ch) && ch != "."
}

func isNumber(ch string) bool {
	return regexp.MustCompile(`\d`).MatchString(ch)
}

func part2(input string) int {
	fmt.Println(input)
	return 0
}
