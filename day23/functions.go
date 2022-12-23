package day23

import (
	"math"
)

func NewArena(input []string) *Arena {
	var a Arena
	a.Elves = []*Elf{}
	a.Props = map[Position]MoveProposal{}
	a.CurrentDir = N
	a.PosLookup = map[Position]bool{}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				a.Elves = append(a.Elves, &Elf{x, y})
				a.PosLookup[Position{x, y}] = true
			}
		}
	}

	return &a
}

func (a *Arena) DoRoundFirstHalf() {
	a.Props = map[Position]MoveProposal{}

	checkOrder := []int{N, S, W, E}
	for i := 0; i < 4; i++ {
		checkOrder[i] = (checkOrder[i] + a.CurrentDir) % 4
	}

	for _, e := range a.Elves {

		hasNeighbor := false
		for dy := -1; dy <= 1 && !hasNeighbor; dy++ {
			for dx := -1; dx <= 1 && !hasNeighbor; dx++ {
				if dy == 0 && dx == 0 {
					continue
				}
				if a.PosLookup[Position{e.X + dx, e.Y + dy}] {
					hasNeighbor = true
				}
			}
		}
		if !hasNeighbor {
			continue
		}

		foundMove := false
		for _, dir := range checkOrder {
			if foundMove {
				break
			}
			switch dir {
			case N:
				if !a.PosLookup[Position{e.X - 1, e.Y - 1}] && !a.PosLookup[Position{e.X, e.Y - 1}] && !a.PosLookup[Position{e.X + 1, e.Y - 1}] {
					ps := Position{e.X, e.Y - 1}
					p := a.Props[ps]
					p.NumProposals++
					p.LastElf = e
					a.Props[ps] = p
					foundMove = true
				}
			case S:
				if !a.PosLookup[Position{e.X - 1, e.Y + 1}] && !a.PosLookup[Position{e.X, e.Y + 1}] && !a.PosLookup[Position{e.X + 1, e.Y + 1}] {
					ps := Position{e.X, e.Y + 1}
					p := a.Props[ps]
					p.NumProposals++
					p.LastElf = e
					a.Props[ps] = p
					foundMove = true
				}
			case W:
				if !a.PosLookup[Position{e.X - 1, e.Y - 1}] && !a.PosLookup[Position{e.X - 1, e.Y}] && !a.PosLookup[Position{e.X - 1, e.Y + 1}] {
					ps := Position{e.X - 1, e.Y}
					p := a.Props[ps]
					p.NumProposals++
					p.LastElf = e
					a.Props[ps] = p
					foundMove = true
				}
			case E:
				if !a.PosLookup[Position{e.X + 1, e.Y - 1}] && !a.PosLookup[Position{e.X + 1, e.Y}] && !a.PosLookup[Position{e.X + 1, e.Y + 1}] {
					ps := Position{e.X + 1, e.Y}
					p := a.Props[ps]
					p.NumProposals++
					p.LastElf = e
					a.Props[ps] = p
					foundMove = true
				}
			}
		}

	}
}

func (a *Arena) DoRoundSecondHalf() bool {
	somebodyMoved := false
	for k, v := range a.Props {
		if v.NumProposals == 1 {
			somebodyMoved = true
			a.PosLookup[Position{v.LastElf.X, v.LastElf.Y}] = false
			v.LastElf.X = k.X
			v.LastElf.Y = k.Y
			a.PosLookup[Position{v.LastElf.X, v.LastElf.Y}] = true
		}
	}
	a.CurrentDir = (a.CurrentDir + 1) % 4
	return somebodyMoved
}

func (a *Arena) DoRound() bool {
	a.DoRoundFirstHalf()
	return a.DoRoundSecondHalf()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a *Arena) GetBounds() Bounds {
	var b Bounds
	b.MinX = math.MaxInt
	b.MinY = math.MaxInt
	b.MaxX = math.MinInt
	b.MaxY = math.MinInt

	for k, v := range a.PosLookup {
		if v {
			b.MinX = min(k.X, b.MinX)
			b.MinY = min(k.Y, b.MinY)
			b.MaxX = max(k.X, b.MaxX)
			b.MaxY = max(k.Y, b.MaxY)
		}
	}
	return b
}

func (a *Arena) countEmptyCells() int {
	bd := a.GetBounds()
	acc := 0
	for y := bd.MinY; y <= bd.MaxY; y++ {
		for x := bd.MinX; x <= bd.MaxX; x++ {
			if !a.PosLookup[Position{x, y}] {
				acc++
			}
		}
	}
	return acc
}
