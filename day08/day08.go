package day08

import (
	"bufio"
	_ "embed"
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
	forest := c.parse()

	maxLeft := make([]rune, len(forest))
	maxRight := make([]rune, len(forest))

	isVisible := make(map[Point]bool)
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if j == 0 {
				maxLeft[i] = forest[i][j]
				maxRight[i] = forest[i][len(forest[0])-1]
				isVisible[Point{i, j}] = true
				isVisible[Point{i, len(forest[0]) - 1}] = true
				continue
			}
			if forest[i][j] > maxLeft[i] {
				isVisible[Point{i, j}] = true
				maxLeft[i] = forest[i][j]
			}
			if forest[i][len(forest[0])-1-j] > maxRight[i] {
				isVisible[Point{i, len(forest[0]) - 1 - j}] = true
				maxRight[i] = forest[i][len(forest[0])-1-j]
			}
		}
	}

	maxTop := make([]rune, len(forest))
	maxDown := make([]rune, len(forest))

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if i == 0 {
				maxTop[j] = forest[i][j]
				maxDown[j] = forest[len(forest)-1][j]
				isVisible[Point{i, j}] = true
				isVisible[Point{len(forest) - 1, j}] = true
				continue
			}
			if forest[i][j] > maxTop[j] {
				isVisible[Point{i, j}] = true
				maxTop[j] = forest[i][j]
			}
			if forest[len(forest)-1-i][j] > maxDown[j] {
				isVisible[Point{len(forest) - 1 - i, j}] = true
				maxDown[j] = forest[len(forest)-1-i][j]
			}
		}
	}

	slog.Info("Part 1", slog.Any("isVisbleCount", len(isVisible)))
}

func (c *Command) part2() {
	forest := c.parse()

	var highestScore int
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			score :=
				forest.CalcViewDown(i, j, forest[i][j], true) *
					forest.CalcViewLeft(i, j, forest[i][j], true) *
					forest.CalcViewRight(i, j, forest[i][j], true) *
					forest.CalcViewTop(i, j, forest[i][j], true)
			if score > highestScore {
				highestScore = score
			}
		}
	}

	slog.Info("Part 2", slog.Any("highestScore", highestScore))
}

func (c *Command) parse() Forest {
	var forest Forest

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		row := []rune{}
		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	return forest
}
