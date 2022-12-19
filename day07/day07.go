package day07

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
}

func (c Command) Execute() {
	c.part1()
	c.part2()
}

func (c Command) part1() {
	fs := c.parse()
	slog.Info("Part 1", slog.Any("totalSize", fs.SumLessThan100000()))
}

func (c Command) part2() {
	fs := c.parse()

	freeSpace := 70000000 - fs.Size
	neededSpace := 30000000 - freeSpace
	result := fs.FindSmallerDir(neededSpace)

	slog.Info("Part 2", slog.Any("result", result))
}

func (c Command) parse() Node {
	fs := NewDir("/", nil)
	fsContext := &fs

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		line := strings.Fields(sc.Text())

		if line[0] == "$" {

			command := line[1]
			if command == "ls" {
				continue
			}

			arg := line[2]

			switch arg {
			case "/":
				fsContext = &fs
			case "..":
				fsContext = fsContext.Parent
			default:
				fsContext = fsContext.Children[arg]
			}
			continue
		}

		if line[0] == "dir" {
			node := NewDir(line[1], fsContext)
			fsContext.Children[node.Name] = &node
		} else {
			size, _ := strconv.Atoi(line[0])
			node := NewFile(line[1], size, fsContext)
			fsContext.Children[node.Name] = &node
		}
	}

	fs.ComputeSize()
	return fs
}
