package day17

import "fmt"

type Chamber struct {
	pieceCounter int
	moveIdx      int
	height       int
	moves        string
	pieces       []Piece
	state        [][7]bool
	cache        map[string]Cache
}

func (c *Chamber) NextPiece() {
	c.pieceCounter++

	pos := [2]int{c.height + 3, 2}
	piece := c.pieces[(c.pieceCounter-1)%len(c.pieces)]

	newLines := (c.height + 7) - len(c.state)
	for i := 0; i < newLines; i++ {
		c.state = append(c.state, [7]bool{})
	}

	for {
		newPos := [2]int{pos[0], pos[1]}
		switch c.moves[c.moveIdx] {
		case '>':
			newPos[1]++
		case '<':
			newPos[1]--
		}

		if !c.isColision(newPos, piece) {
			pos = newPos
		}

		c.moveIdx = (c.moveIdx + 1) % len(c.moves)

		if c.isColision([2]int{pos[0] - 1, pos[1]}, piece) {
			c.savePiece(pos, piece)
			return
		}
		pos[0]--

	}
}

func (c *Chamber) UpdateSeenState() Cache {
	if len(c.state) < 30 {
		return Cache{}
	}

	pieceIdx := c.pieceCounter % len(c.pieces)
	stateHash := fmt.Sprintf("%03d:%05d:", pieceIdx, c.moveIdx)
	for l := 0; l < 30; l++ {
		lineHash := 0
		for x := 0; x < 7; x++ {
			if c.state[l][6-x] {
				lineHash |= 1 << x
			}
		}
		stateHash += fmt.Sprintf("%03d", lineHash)
	}

	if x, ok := c.cache[stateHash]; ok {
		return x
	}

	c.cache[stateHash] = Cache{
		height:     c.height,
		pieceCount: c.pieceCounter,
	}
	return Cache{}
}

func (c *Chamber) isColision(pos [2]int, piece Piece) bool {
	if pos[0] < 0 || pos[1] < 0 || pos[1]+piece.width > len(c.state[0]) {
		return true
	}

	for line := 0; line < len(piece.shape); line++ {
		for col := 0; col < len(piece.shape[line]); col++ {
			if piece.shape[line][col] && c.state[pos[0]+line][pos[1]+col] {
				return true
			}
		}
	}

	return false
}

func (c *Chamber) savePiece(pos [2]int, piece Piece) {
	for line := 0; line < len(piece.shape); line++ {
		for col := 0; col < len(piece.shape[line]); col++ {
			if piece.shape[line][col] {
				c.state[pos[0]+line][pos[1]+col] = true
				if pos[0]+line+1 > c.height {
					c.height = pos[0] + line + 1
				}
			}
		}
	}
}
