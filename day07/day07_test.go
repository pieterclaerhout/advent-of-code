package day07_test

import (
	"testing"

	"github.com/pieterclaerhout/advent-of-code/day07"
	"github.com/stretchr/testify/assert"
)

const input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestDay4(t *testing.T) {
	cmd := day07.Command{}
	result1, result2 := cmd.Execute(input)

	assert.Equal(t, 95437, result1)
	assert.Equal(t, 24933642, result2)
}
