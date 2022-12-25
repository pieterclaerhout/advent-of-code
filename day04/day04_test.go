package day04_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day04"
	"github.com/stretchr/testify/assert"
)

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestDay4(t *testing.T) {
	cmd := day04.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 2, result1)
	assert.Equal(t, 4, result2)
}
