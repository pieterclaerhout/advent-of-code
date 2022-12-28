package day03_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day03"
	"github.com/stretchr/testify/assert"
)

const input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestDay3(t *testing.T) {
	cmd := day03.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 157, result1)
	assert.Equal(t, 70, result2)
}
