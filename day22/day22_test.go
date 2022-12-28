package day22_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day22"
	"github.com/stretchr/testify/assert"
)

const input = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

func Test_Command(t *testing.T) {
	cmd := day22.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 6032, result1)
	assert.Equal(t, 5031, result2)
}
