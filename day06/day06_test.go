package day06_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day06"
	"github.com/stretchr/testify/assert"
)

const input = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func TestDay4(t *testing.T) {
	cmd := day06.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 7, result1)
	assert.Equal(t, 19, result2)
}
