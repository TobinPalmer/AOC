package main

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"os"
	"sort"
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

var types = map[string]uint8{
	"five":       0,
	"four":       1,
	"full_house": 2,
	"three_kind": 3,
	"two_pair":   4,
	"one_pair":   5,
	"high_card":  6,
}

func part1(input string) int {
	seniority := map[string]uint8{
		"A": 0,
		"K": 1,
		"Q": 2,
		"J": 3,
		"T": 4,
		"9": 5,
		"8": 6,
		"7": 7,
		"6": 8,
		"5": 9,
		"4": 10,
		"3": 11,
		"2": 12,
	}
	split := strings.Split(input, "\n")
	var handMap []interface{}
	for _, line := range split {
		input := line[:5]
		bid := line[6:]
		val, _ := strconv.Atoi(bid)
		handMap = append(handMap, []interface{}{input, val})
	}

	handArr := make([]string, 0, len(handMap))

	for _, v := range handMap {
		innerslice := v.([]interface{})
		handArr = append(handArr, innerslice[0].(string))
	}

	sort.Slice(handArr, func(i, j int) bool {
		for idx := 0; idx < len(handArr[i]); idx++ {
			if handArr[i][idx] == handArr[j][idx] {
				continue
			}

			if getType(handArr[j]) == getType(handArr[i]) {
				if seniority[string(handArr[i][idx])] < seniority[string(handArr[j][idx])] {
					return true
				} else {
					return false
				}
			} else {
				return getType(handArr[j]) > getType(handArr[i])
			}

		}

		return false
	})

	var bestMap []interface{}

	total := 0

outer:
	for i := 0; i < len(handArr); i++ {
		for j := 0; j < len(handMap); j++ {
			if handArr[i] == handMap[j].([]interface{})[0].(string) {
				bestMap = append(bestMap, []interface{}{handArr[i], handMap[j].([]interface{})[1].(int)})
				continue outer
			}
		}
	}

	for i, j := len(handArr)-1, 1; i >= 0; i, j = i-1, j+1 {
		innerslice := bestMap[i].([]interface{})
		total += innerslice[1].(int) * j
	}

	return total
}

func getType(card string) uint8 {
	letters := make(map[uint8]uint8)

	for i := 0; i < len(card); i++ {
		if val, ok := letters[card[i]]; ok {
			letters[card[i]] = val + 1
		} else {

			letters[card[i]] = 1
		}
	}

	if len(letters) == 1 {
		return types["five"]
	}

	if len(letters) == 5 {
		return types["high_card"]
	}

	var maxLetters uint8 = 0

	for _, v := range letters {
		maxLetters = max(maxLetters, v)
	}

	if maxLetters == 3 && len(letters) == 3 {
		return types["three_kind"]
	}

	if maxLetters == 2 && len(letters) == 3 {
		return types["two_pair"]
	}

	if maxLetters == 4 && len(letters) == 2 {
		return types["four"]
	}

	if maxLetters == 3 && len(letters) == 3 {
		return types["full_house"]
	}

	if maxLetters == 2 && len(letters) == 4 {
		return types["one_pair"]
	}

	if maxLetters == 3 && len(letters) == 2 {
		return types["full_house"]
	}

	fmt.Println("This shouldn't print", card)

	return 100
}

func part2(input string) int {
	split := strings.Split(input, "\n")

	for i := 0; i < len(split); i++ {
	}

	return 0
}
