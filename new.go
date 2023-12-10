package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func main() {
	const year = 2023
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide the day number, ex 5")
		return
	}
	file, err := create("2023/day" + args[0] + "/main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	defaultContent := []byte(`
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

	for i := 0; i < len(split); i++ {

	}

	return 0
}

func part2(input string) int {
	split := strings.Split(input, "\n")

	for i := 0; i < len(split); i++ {
	}

	return 0
}
`)
	_, err = file.Write(defaultContent)
	if err != nil {
		fmt.Println("Error writing to main file")
		return
	}

	file, err = create("2023/day" + args[0] + "/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	defaultTestContent := []byte(`
package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: "",
			want:  0,
		},
		//{
		//	name:  "actual",
		//	input: input,
		//	want:  0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example",
			input: "",
			want: 0,
		},
		//{
		//	name:  "actual",
		//	input: input,
		//	want:  0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
`)
	file, err = create("2023/day" + args[0] + "/main_test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.Write(defaultTestContent)
	if err != nil {
		fmt.Println("Error writing to testing file")
		return
	}
}
