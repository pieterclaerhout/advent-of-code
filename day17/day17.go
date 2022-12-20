package day17

import (
	_ "embed"
	"strings"

	"golang.org/x/exp/slog"
)

//go:embed input.txt
var input string

type Command struct {
	RegisterX   int
	CycleNumber int
	FinalValue  int
}

func (c *Command) Execute() {
	c.part1()
	c.part2()
}

func (c *Command) part1() {
	chamber := Chamber{
		moves:  strings.TrimSpace(input),
		pieces: Pieces,
		state:  [][7]bool{},
	}

	for i := 0; i < 2022; i++ {
		chamber.NextPiece()
	}

	simulation := Simulation{}
	height := simulation.runSimulation(input, 2022) + 1

	slog.Info("Part 1", slog.Any("result", chamber.height), slog.Any("result2", height+1))
}

func (c *Command) part2() {
	// chamber := Chamber{
	// 	moves:  strings.TrimSpace(input),
	// 	pieces: Pieces,
	// 	state:  [][7]bool{},
	// 	cache:  map[string]Cache{},
	// }

	// for i := 0; i < 1000000000000; i++ {
	// 	chamber.NextPiece()

	// 	cacheState := chamber.UpdateSeenState()
	// 	if cacheState.height > 0 {
	// 		remaining := 1000000000000 - chamber.pieceCounter
	// 		repeatedHeight := chamber.height - cacheState.height
	// 		repeatedLen := chamber.pieceCounter - cacheState.pieceCount
	// 		repeatedTotalHeight := repeatedHeight * (remaining / repeatedLen)

	// 		for j := 0; j < remaining%repeatedLen; j++ {
	// 			chamber.NextPiece()
	// 		}

	// 		chamber.height += repeatedTotalHeight
	// 		break
	// 	}
	// }

	// height := c.Calculate(1000000000000)

	simulation := Simulation{}
	height := simulation.runSimulation2(input, 1_000_000_000_000) + 1
	slog.Info("Part 2", slog.Any("result", height))
}

// func (c *Command) Calculate(n int) int {
// 	height := 0
// 	occ := map[point]bool{}
// 	cache := map[string]struct {
// 		round  int
// 		height int
// 	}{}
// 	contentIdx := 0

// 	for round := 1; round < n; round++ {
// 		r := rocks[(round-1)%len(rocks)]

// 		newRock := make(rock, len(r))
// 		for i := 0; i < len(r); i++ {
// 			newRock[i].x = r[i].x
// 			newRock[i].y = r[i].y + height + 3
// 		}

// 		for {
// 			if contentIdx >= len(input) {
// 				contentIdx = 0
// 			}
// 			direction := input[contentIdx]
// 			contentIdx++

// 			var rockTemp rock
// 			for _, dir := range []point{directions[direction], directions['v']} {
// 				rockTemp = make(rock, len(newRock))
// 				for i, p := range newRock {
// 					p.x += dir.x
// 					p.y += dir.y

// 					if p.x < 0 || p.x > 6 || p.y < 0 {
// 						rockTemp = nil
// 						break
// 					}
// 					if _, ok := occ[p]; ok {
// 						rockTemp = nil
// 						break
// 					}
// 					rockTemp[i] = p
// 				}

// 				if rockTemp != nil {
// 					newRock = rockTemp
// 				}
// 			}

// 			if rockTemp != nil {
// 				continue
// 			}

// 			for _, p := range newRock {
// 				occ[p] = true
// 				if p.y+1 > height {
// 					height = p.y + 1
// 				}
// 			}

// 			if round == 2022 {
// 				fmt.Println("Part 1: ", height)
// 			}

// 			key := fmt.Sprintf("%d%d", (round-1)%len(rocks), contentIdx-1)
// 			if val, ok := cache[key]; ok {
// 				quotient := (n - round) / (round - val.round)
// 				remainder := (n - round) % (round - val.round)

// 				if remainder == 0 {
// 					return height + (height-val.height)*quotient
// 				}
// 			} else {
// 				cache[key] = struct {
// 					round  int
// 					height int
// 				}{round: round, height: height}
// 			}
// 			break
// 		}
// 	}

// 	return 0
// }
