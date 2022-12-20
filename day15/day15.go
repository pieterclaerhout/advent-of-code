package day15

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
	RegisterX   int
	CycleNumber int
	FinalValue  int
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	sensors := c.parse()
	slog.Info("Part 1", slog.Any("result", sensors.Coverage(2000000)))
}

func (c *Command) part2() {
	sensors := c.parse()
	slog.Info("Part 2", slog.Any("result", sensors.Coverage2D(4000000)))
}

func (c *Command) parse() Sensors {
	sensors := Sensors{}

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		sensor := Sensor{
			Position: Point{},
			Beacon:   Point{},
		}

		fmt.Sscanf(
			sc.Text(),
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.Position.X,
			&sensor.Position.Y,
			&sensor.Beacon.X,
			&sensor.Beacon.Y,
		)

		sensors = append(sensors, sensor)
	}

	return sensors
}
