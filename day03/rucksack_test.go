package day03_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day03"
	"github.com/stretchr/testify/assert"
)

func Test_Score(t *testing.T) {
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	ruckSacks := []day03.RuckSack{}

	for _, i := range input {
		ruckSacks = append(ruckSacks, day03.NewRuckSack(i))
	}

	var sum int
	for _, ruckSack := range ruckSacks {
		t.Log(ruckSack.Common(), ruckSack.Priorities())
		sum += ruckSack.Priorities()
	}

	assert.Equal(t, 157, sum)
}
