package day17

import (
	"io"
	"math/bits"
)

type Command struct{}

func (cmd *Command) Execute(input string) (any, any) {
	return sim(input, 2022), sim(input, 1e12)
}

type shape []byte

func (s shape) coll(cb []byte, dx, dy int) bool {
	if dx < 0 || dx > 7-s.width() {
		return true
	}
	for i := 0; i < len(s) && dy+i < len(cb); i++ {
		if cb[dy+i]&(s[i]>>dx) != 0 {
			return true
		}
	}
	return false
}

func (s shape) land(cb []byte, dx, dy int) []byte {
	if len(cb) < dy+len(s) {
		cb = append(cb, make([]byte, dy+len(s)-len(cb))...)
	}
	for i, b := range s {
		cb[dy+i] |= b >> dx
	}
	return cb
}

func (s shape) width() (n int) {
	for _, b := range s {
		if n2 := 8 - bits.TrailingZeros8(b); n2 > n {
			n = n2
		}
	}
	return n
}

var (
	// ####
	shp0 shape = []byte{0b1111_0000}

	// .#.
	// ###
	// .#.
	shp1 shape = []byte{0b0100_0000, 0b1110_0000, 0b0100_0000}

	// ..#
	// ..#
	// ###
	shp2 shape = []byte{0b1110_0000, 0b0010_0000, 0b0010_0000}

	// #
	// #
	// #
	// #
	shp3 shape = []byte{0b1000_0000, 0b1000_0000, 0b1000_0000, 0b1000_0000}

	// ##
	// ##
	shp4 shape = []byte{0b1100_0000, 0b1100_0000}
)

func sim(input string, n int) int {
	js := []byte(input)
	cb := []byte{0b1111_1111}
	ss := [...]shape{shp0, shp1, shp2, shp3, shp4}
	for i, j, co := 0, 0, []int{}; i < n; i++ {
		for dx, dy, s := 2, len(cb)+3, ss[i%len(ss)]; ; dy, j = dy-1, j+1 {
			if s.coll(cb, dx, dy) {
				cb = s.land(cb, dx, dy+1)
				break
			}
			if js[j%len(js)] == '<' && !s.coll(cb, dx-1, dy) {
				dx--
			}
			if js[j%len(js)] == '>' && !s.coll(cb, dx+1, dy) {
				dx++
			}
		}
		if j < len(js)*2 /* warmup */ {
			continue
		}
		if len(co) == 0 /* cutoff */ {
			co = append(co, i, j%len(js), digest(cb), len(cb))
			continue
		}
		if o := i - co[0]; o%len(ss) == 0 && j%len(js) == co[1] && digest(cb) == co[2] {
			return (n-co[0])/o*(len(cb)-co[3]) + co[(n-co[0])%o+2] - 1
		}
		co = append(co, len(cb))
	}
	return len(cb) - 1
}

func digest(bs []byte) (n int) {
	for i := 0; i < 9; i++ {
		n = n<<7 | (int(bs[i]) & 0x7f)
	}
	return n
}

func scan(r io.Reader) (ps []byte) {
	ps, _ = io.ReadAll(r)
	return ps
}
