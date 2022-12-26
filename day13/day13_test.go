package day13_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day13"
	"github.com/stretchr/testify/assert"
)

const input = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func TestDay4(t *testing.T) {
	cmd := day13.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 13, result1)
	assert.Equal(t, 140, result2)
}
