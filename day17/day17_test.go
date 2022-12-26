package day17_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day17"
	"github.com/stretchr/testify/assert"
)

const input = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestDay4(t *testing.T) {
	cmd := day17.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 3068, result1)
	assert.Equal(t, 1514285714288, result2)
}
