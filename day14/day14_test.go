package day14_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day14"
	"github.com/stretchr/testify/assert"
)

const input = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestDay4(t *testing.T) {
	cmd := day14.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 24, result1)
	assert.Equal(t, 93, result2)
}
