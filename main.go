package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pieterclaerhout/advent-of-code/day01"
	"github.com/pieterclaerhout/advent-of-code/day02"
	"github.com/pieterclaerhout/advent-of-code/day03"
	"github.com/pieterclaerhout/advent-of-code/day04"
	"github.com/pieterclaerhout/advent-of-code/day05"
	"github.com/pieterclaerhout/advent-of-code/day06"
	"github.com/pieterclaerhout/advent-of-code/day07"
	"github.com/pieterclaerhout/advent-of-code/day08"
	"github.com/pieterclaerhout/advent-of-code/day09"
	"github.com/pieterclaerhout/advent-of-code/day10"
	"github.com/pieterclaerhout/advent-of-code/day11"
	"github.com/pieterclaerhout/advent-of-code/day12"
	"github.com/pieterclaerhout/advent-of-code/day13"
	"github.com/pieterclaerhout/advent-of-code/day14"
	"github.com/pieterclaerhout/advent-of-code/day15"
	"github.com/pieterclaerhout/advent-of-code/day16"
	"github.com/pieterclaerhout/advent-of-code/day17"
	"github.com/pieterclaerhout/advent-of-code/day18"
	"github.com/pieterclaerhout/advent-of-code/day20"
)

var day = flag.Int("day", 0, "day to execute")

type Command interface {
	Execute(input string) (any, any)
}

func main() {
	flag.Parse()

	commands := []Command{
		&day01.Command{},
		&day02.Command{},
		&day03.Command{},
		&day04.Command{},
		&day05.Command{},
		&day06.Command{},
		&day07.Command{},
		&day08.Command{},
		&day09.Command{},
		&day10.Command{},
		&day11.Command{},
		&day12.Command{},
		&day13.Command{},
		&day14.Command{},
		&day15.Command{Row1: 2000000, Row2: 4000000},
		&day16.Command{},
		&day17.Command{},
		&day18.Command{},
		// &day19.Command{},
		&day20.Command{},
		// &day21.Command{},
		// &day22.Command{},
		// &day23.Command{},
		// &day24.Command{},
		// &day25.Command{},
	}

	if *day == 0 {
		flag.PrintDefaults()
		return
	}

	if *day > len(commands) {
		fmt.Println("Command not found: day", *day)
		return
	}

	command := commands[*day-1]

	inputPath := filepath.Join(fmt.Sprintf("day%02d", *day), "input.txt")
	rawInput, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Failed to read:", inputPath)
		return
	}

	input := string(rawInput)
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")
	input = strings.TrimRight(input, "\n")

	result1, result2 := command.Execute(string(input))

	fmt.Println("Part 1:", result1)
	fmt.Println("Part 2:", result2)
}
