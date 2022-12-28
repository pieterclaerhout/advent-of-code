package day21_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day21"
	"github.com/stretchr/testify/assert"
)

const input = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func Test_Command(t *testing.T) {
	cmd := day21.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 152, result1)
	assert.Equal(t, 301, result2)
}
