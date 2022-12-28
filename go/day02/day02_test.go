package day02_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day02"
	"github.com/stretchr/testify/assert"
)

const input = `A Y
B X
C Z`

func TestDay2(t *testing.T) {
	cmd := day02.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 15, result1)
	assert.Equal(t, 12, result2)
}
