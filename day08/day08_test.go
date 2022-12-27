package day08_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day08"
	"github.com/stretchr/testify/assert"
)

const input = `30373
25512
65332
33549
35390`

func Test_Command(t *testing.T) {
	cmd := day08.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 21, result1)
	assert.Equal(t, 8, result2)
}
