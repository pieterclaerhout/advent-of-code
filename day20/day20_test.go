package day20_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day20"
	"github.com/stretchr/testify/assert"
)

func Test_Mix(t *testing.T) {
	input := []int{1, 2, -3, 3, -2, 0, 4}

	front := day20.NewList(input)
	day20.Mix(front, len(input))
	sum := day20.GroveCoordinates(front)

	assert.Equal(t, 3, sum)
}

func Test_DecryptKey(t *testing.T) {
	input := []int{1, 2, -3, 3, -2, 0, 4}

	front := day20.NewList(input)
	day20.ApplyDecryptKey(front, 811589153)

	for i := 0; i < 10; i++ {
		day20.Mix(front, len(input))
	}

	sum := day20.GroveCoordinates(front)

	assert.Equal(t, 1623178306, sum)
}
