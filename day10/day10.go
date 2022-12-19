package day10

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
	RegisterX   int
	CycleNumber int
	FinalValue  int
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	c.CycleNumber = 0
	c.RegisterX = 1
	c.FinalValue = 0

	for _, operation := range c.parse() {
		c.incrementAndControl1()
		if operation.Type == "addx" {
			c.incrementAndControl1()
			c.RegisterX += operation.Value
		}
	}

	slog.Info("Part 1", slog.Any("finalValue", c.FinalValue))
}

func (c *Command) part2() {
	c.CycleNumber = 0
	c.RegisterX = 1
	c.FinalValue = 0

	for _, operation := range c.parse() {
		c.incrementAndControl2()
		if operation.Type == "addx" {
			c.incrementAndControl2()
			c.RegisterX += operation.Value
		}
	}

	// slog.Info("Part 1", slog.Any("finalValue", c.FinalValue))
}

func (c *Command) parse() []Operation {
	var list []Operation

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		line := strings.Fields(sc.Text())

		operation := line[0]
		var value int
		if len(line) > 1 {
			value, _ = strconv.Atoi(line[1])
		}

		list = append(list, Operation{
			Type:  operation,
			Value: value,
		})
	}

	return list
}

func (c *Command) incrementAndControl1() {
	c.CycleNumber++
	if (c.CycleNumber-20)%40 == 0 && c.CycleNumber <= 220 {
		c.FinalValue += c.RegisterX * c.CycleNumber
	}
}

func (c *Command) incrementAndControl2() {
	if c.CycleNumber%40 == 0 && c.CycleNumber <= 220 {
		fmt.Println()
	}
	if c.RegisterX-1 == c.CycleNumber%40 || c.RegisterX == c.CycleNumber%40 || c.RegisterX+1 == c.CycleNumber%40 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	c.CycleNumber++
}
