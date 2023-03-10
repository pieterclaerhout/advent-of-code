package day01

import (
	"sort"
	"strconv"
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	calories := cmd.parse((input))

	max := calories[0]
	top3 := cmd.sum(calories, 3)

	return max, top3
}

func (cmd *Command) parse(input string) []int {
	result := []int{}

	for _, chunk := range strings.Split(input, "\n\n") {
		chunkSum := 0

		for _, line := range strings.Split(chunk, "\n") {
			value, _ := strconv.Atoi(line)
			chunkSum += value
		}

		result = append(result, chunkSum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result
}

func (cmd *Command) sum(ints []int, count int) int {
	sum := 0
	for i := 0; i < len(ints) && i < count; i++ {
		sum += ints[i]
	}
	return sum
}
