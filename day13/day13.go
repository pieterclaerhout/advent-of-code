package day13

import (
	"bufio"
	_ "embed"
	"sort"
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
	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	index := 1
	var indexSum int

	for sc.Scan() {
		package1 := c.readTree(sc.Text())
		sc.Scan()
		package2 := c.readTree(sc.Text())

		if c.areOrdered(package1, package2) == 1 {
			indexSum += index
		}

		index++
		sc.Scan()
	}

	slog.Info("Part 1", slog.Any("result", indexSum))
}

func (c *Command) part2() {
	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	var packages []Tree
	for sc.Scan() {
		packages = append(packages, c.readTree(sc.Text()))
		sc.Scan()
		packages = append(packages, c.readTree(sc.Text()))
		sc.Scan()
	}
	packages = append(packages, c.readTree("[[2]]"))
	packages = append(packages, c.readTree("[[6]]"))

	sort.Slice(packages, func(i, j int) bool {
		return c.areOrdered(packages[i], packages[j]) == 1
	})

	decoderKey := 1
	for i, p := range packages {
		if c.areOrdered(p, c.readTree("[[2]]")) == 0 || c.areOrdered(p, c.readTree("[[6]]")) == 0 {
			decoderKey *= i + 1
		}
	}

	slog.Info("Part 2", slog.Any("result", decoderKey))
}

func (c *Command) readTree(input string) Tree {
	root := NewTree(nil)
	temp := &root

	var currentNumber string
	for _, r := range input {
		switch r {
		case '[':
			newTree := NewTree(temp)
			temp.Elements = append(temp.Elements, &newTree)
			temp = &newTree
		case ']':
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.ValueLeaf = number
				currentNumber = ""
			}
			temp = temp.Father
		case ',':
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.ValueLeaf = number
				currentNumber = ""
			}
			temp = temp.Father
			newTree := NewTree(temp)
			temp.Elements = append(temp.Elements, &newTree)
			temp = &newTree
		default:
			currentNumber += string(r)
		}
	}

	return root
}

func (c *Command) areOrdered(first Tree, second Tree) int {
	switch {
	case len(first.Elements) == 0 && len(second.Elements) == 0:
		if first.ValueLeaf > second.ValueLeaf {
			return -1
		} else if first.ValueLeaf == second.ValueLeaf {
			return 0
		}
		return 1

	case first.ValueLeaf >= 0:
		return c.areOrdered(Tree{-1, []*Tree{&first}, nil}, second)

	case second.ValueLeaf >= 0:
		return c.areOrdered(first, Tree{-1, []*Tree{&second}, nil})

	default:
		var i int
		for i = 0; i < len(first.Elements) && i < len(second.Elements); i++ {
			ordered := c.areOrdered(*first.Elements[i], *second.Elements[i])
			if ordered != 0 {
				return ordered
			}
		}
		if i < len(first.Elements) {
			return -1
		} else if i < len(second.Elements) {
			return 1
		}
	}
	return 0
}
