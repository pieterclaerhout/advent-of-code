package day03

import (
	"strings"
)

type RuckSack struct {
	Compartment1 []string
	Compartment2 []string
}

func (rs RuckSack) Priorities() int {
	var priorities int
	for _, c := range rs.Common() {
		priorities += scoreForCharacter(c)
	}
	return priorities
}

func (rs RuckSack) Common() []string {
	common := intersect(rs.Compartment1, rs.Compartment2)
	return common[:1]
}

func (rs RuckSack) All() []string {
	return append(rs.Compartment1, rs.Compartment2...)
}

func NewRuckSack(input string) RuckSack {
	return RuckSack{
		Compartment1: strings.Split(input[:len(input)/2], ""),
		Compartment2: strings.Split(input[len(input)/2:], ""),
	}
}
