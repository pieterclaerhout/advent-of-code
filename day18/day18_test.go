package day18_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day18"
	"github.com/stretchr/testify/assert"
)

const input = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestDay4(t *testing.T) {
	cmd := day18.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 64, result1)
	assert.Equal(t, 58, result2)
}
