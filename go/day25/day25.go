package day25

import (
	"strings"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	sum := cmd.snafuToDecimal(input)
	snafu := cmd.decimalToSnafu(sum)

	return sum, snafu
}

func (cmd *Command) snafuToDecimal(input string) int {
	mapping := map[rune]int{
		'=': -2,
		'-': -1,
		'0': 0,
		'1': 1,
		'2': 2,
	}

	sum := 0
	for _, s := range strings.Fields(input) {
		n := 0
		for _, r := range s {
			n = 5*n + mapping[r]
		}
		sum += n
	}

	return sum
}

func (cmd *Command) decimalToSnafu(sum int) string {
	mapping := []string{"=", "-", "0", "1", "2"}

	snafu := ""
	for sum > 0 {
		snafu = mapping[(sum+2)%5] + snafu
		sum = (sum + 2) / 5
	}

	return snafu
}
