package day25_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day25"
	"github.com/stretchr/testify/assert"
)

const input = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func Test_Command(t *testing.T) {
	cmd := day25.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 4890, result1)
	assert.Equal(t, "2=-1=0", result2)
}
