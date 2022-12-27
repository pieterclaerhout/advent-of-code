package day20_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day20"
	"github.com/stretchr/testify/assert"
)

const input = `1
2
-3
3
-2
0
4`

func TestDay4(t *testing.T) {
	cmd := day20.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 3, result1)
	assert.Equal(t, 1623178306, result2)
}
