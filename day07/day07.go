package day07

import (
	"strconv"
	"strings"
)

type Command struct {
}

func (c *Command) Execute(input string) (interface{}, interface{}) {
	fs := c.parse(input)

	return c.part1(fs), c.part2(fs)
}

func (c *Command) part1(fs Node) int {
	return fs.SumLessThan100000()
}

func (c *Command) part2(fs Node) int {
	freeSpace := 70000000 - fs.Size
	neededSpace := 30000000 - freeSpace

	return fs.FindSmallerDir(neededSpace)
}

func (c *Command) parse(input string) Node {
	fs := NewDir("/", nil)
	fsContext := &fs

	for _, line := range strings.Split(input, "\n") {
		lineParts := strings.Fields(line)

		if lineParts[0] == "$" {

			command := lineParts[1]
			if command == "ls" {
				continue
			}

			arg := lineParts[2]

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

		if lineParts[0] == "dir" {
			node := NewDir(lineParts[1], fsContext)
			fsContext.Children[node.Name] = &node
		} else {
			size, _ := strconv.Atoi(lineParts[0])
			node := NewFile(lineParts[1], size, fsContext)
			fsContext.Children[node.Name] = &node
		}
	}

	fs.ComputeSize()
	return fs
}
