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

	times := stringArrToIntArr(strings.Split(strings.TrimSpace(split[0][strings.Index(split[0], ":")+1:]), " "))
	distances := stringArrToIntArr(strings.Split(strings.TrimSpace(split[1][strings.Index(split[1], ":")+1:]), " "))
	timesToDistances := make(map[int]int)

	if len(times) != len(distances) {
		panic("Expected time and distance ratio to be 1:1")
	}

	for i := 0; i < len(times); i++ {
		timesToDistances[times[i]] = distances[i]
	}

	total := 0

	for input, prevMax := range timesToDistances {
		tick := 0
		currSpeed := 0
		currDistance := 0
		moreThanMax := 0

		for j := 0; j < input+1; j++ {
			for k := j; k > 0; k-- {
				tick++
				currSpeed++
			}

			for remainingTicks := input - tick; remainingTicks > 0; remainingTicks-- {
				currDistance += currSpeed
			}

			if currDistance > prevMax {
				moreThanMax++
			}

			tick = 0
			currSpeed = 0
			currDistance = 0
		}

		if total == 0 {
			total = moreThanMax
		} else {
			total *= moreThanMax
		}
	}

	return total
}

func stringArrToIntArr(input []string) []int {
	var output []int

	for i := 0; i < len(input); i++ {
		val, err := strconv.Atoi(input[i])
		if err != nil {
			continue
		}

		output = append(output, val)
	}

	return output
}

func part2(input string) int {
	split := strings.Split(input, "\n")

	distances, _ := strconv.ParseInt(strings.ReplaceAll(split[1][strings.Index(split[1], ":")+1:], " ", ""), 10, 64)
	time, _ := strconv.ParseInt(strings.ReplaceAll(split[0][strings.Index(split[0], ":")+1:], " ", ""), 10, 64)

	var fi int64 = 0
	var l int64 = 0

	for i := int64(0); i < time; i++ {
		if f(time, i, distances) == true {
			fi = i
			break
		}
	}

	for i := time; i > 0; i-- {
		if f(time, i, distances) == true {
			l = i
			break
		}
	}

	return int(l - fi + 1)
}

func f(time int64, tick int64, d int64) bool {
	return (tick * (time - tick)) > d
}
