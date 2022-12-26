package day12_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day12"
	"github.com/stretchr/testify/assert"
)

const input = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestDay4(t *testing.T) {
	cmd := day12.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 31, result1)
	assert.Equal(t, 29, result2)
}
