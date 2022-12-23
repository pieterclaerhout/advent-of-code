package day23

type Elf struct {
	X int
	Y int
}

type Position struct {
	X int
	Y int
}

type MoveProposal struct {
	NumProposals int
	LastElf      *Elf
}

type Bounds struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

type Arena struct {
	Elves      []*Elf
	Props      map[Position]MoveProposal
	CurrentDir int
	PosLookup  map[Position]bool
}
