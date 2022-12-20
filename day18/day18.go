package day18

import (
	"bufio"
	_ "embed"
	"fmt"
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
	cubes := c.parse()
	_, area := cubes.BuildMap()
	slog.Info("Part 1", slog.Any("result", area))
}

func (c *Command) part2() {
	cubes := c.parse()
	cubesMap := cubes.BuildMap3D()
	cubesMap.DfsSearch(Cube{0, 0, 0})
	area := cubesMap.Area()
	slog.Info("Part 2", slog.Any("result", area))
}

func (c *Command) parse() Cubes {
	cubes := Cubes{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		cube := Cube{}
		fmt.Sscanf(sc.Text(), "%d,%d,%d",
			&cube.X,
			&cube.Y,
			&cube.Z,
		)
		cubes = append(cubes, cube)
	}

	return cubes
}
