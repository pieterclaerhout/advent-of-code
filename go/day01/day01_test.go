package day01_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day01"
	"github.com/stretchr/testify/assert"
)

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestDay1(t *testing.T) {
	cmd := day01.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 24000, result1)
	assert.Equal(t, 45000, result2)
}
