package day11

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	monkeys := c.parse(3)
	counts := make([]int, len(monkeys))

	for turn := 0; turn < 20; turn++ {
		for monkeyId, currentMonkey := range monkeys {
			for _, item := range currentMonkey.Items {
				newValue := currentMonkey.Operation(item)
				monkeys[currentMonkey.TestAndThrow(newValue)].Items = append(monkeys[currentMonkey.TestAndThrow(newValue)].Items, newValue)
			}
			counts[monkeyId] += len(monkeys[monkeyId].Items)
			monkeys[monkeyId].Items = []int{}
		}
	}

	var highestCount int
	var secondHighest int
	for _, count := range counts {
		if count > secondHighest {
			secondHighest = count
		}
		if secondHighest > highestCount {
			highestCount, secondHighest = secondHighest, highestCount
		}
	}

	slog.Info("Part 1", slog.Any("result", highestCount*secondHighest))
}

func (c *Command) part2() {
	monkeys := c.parse(1)
	counts := make([]int, len(monkeys))

	var bigLimit int = 1
	for _, monkey := range monkeys {
		bigLimit *= monkey.TestingValue
	}

	for turn := 0; turn < 10000; turn++ {
		for monkeyId, currentMonkey := range monkeys {
			for _, item := range currentMonkey.Items {
				newValue := currentMonkey.Operation(item) % bigLimit
				monkeys[currentMonkey.TestAndThrow(newValue)].Items = append(monkeys[currentMonkey.TestAndThrow(newValue)].Items, newValue)
			}
			counts[monkeyId] += len(monkeys[monkeyId].Items)
			monkeys[monkeyId].Items = []int{}
		}
	}

	var highestCount int
	var secondHighest int
	for _, count := range counts {
		if count > secondHighest {
			secondHighest = count
		}
		if secondHighest > highestCount {
			highestCount, secondHighest = secondHighest, highestCount
		}
	}

	slog.Info("Part 2", slog.Any("result", highestCount*secondHighest))
}

func (c *Command) parse(divider int) []Monkey {
	var list []Monkey

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		sc.Scan()

		var newMonkey Monkey
		for _, item := range strings.Split(sc.Text()[len("  Starting items: "):], ", ") {
			worryLevel, _ := strconv.Atoi(item)
			newMonkey.Items = append(newMonkey.Items, worryLevel)
		}

		sc.Scan()
		var operator rune
		var value int = 0
		if sc.Text() == "  Operation: new = old * old" || sc.Text() == "  Operation: new = old + old" {
			fmt.Sscanf(sc.Text(), "  Operation: new = old %c old", &operator)
		} else {
			fmt.Sscanf(sc.Text(), "  Operation: new = old %c %d", &operator, &value)
		}
		newMonkey.Operation = c.createOperation(operator, value, divider)

		sc.Scan()
		fmt.Sscanf(sc.Text(), "  Test: divisible by %d", &newMonkey.TestingValue)

		sc.Scan()
		var toThrowIfTrue int
		fmt.Sscanf(sc.Text(), "    If true: throw to monkey %d", &toThrowIfTrue)

		sc.Scan()
		var toThrowIfFalse int
		fmt.Sscanf(sc.Text(), "    If false: throw to monkey %d", &toThrowIfFalse)
		newMonkey.TestAndThrow = c.createTestAndThrow(newMonkey.TestingValue, toThrowIfTrue, toThrowIfFalse)

		sc.Scan()
		list = append(list, newMonkey)
	}

	return list
}

func (c *Command) createOperation(operator rune, value int, divider int) func(int) int {
	return func(n int) int {
		var valueToUse int = value
		if valueToUse == 0 {
			valueToUse = n
		}
		if operator == '+' {
			return (n + valueToUse) / divider
		}
		return (n * valueToUse) / divider
	}
}

func (c *Command) createTestAndThrow(testingValue, toThrowIfTrue, toThrowIfFalse int) func(int) int {
	return func(n int) int {
		if n%testingValue == 0 {
			return toThrowIfTrue
		}
		return toThrowIfFalse
	}
}
