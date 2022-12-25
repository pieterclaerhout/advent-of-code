package day09_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day09"
	"github.com/stretchr/testify/assert"
)

const input = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestDay4(t *testing.T) {
	cmd := day09.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 13, result1)
	assert.Equal(t, 1, result2)
}
