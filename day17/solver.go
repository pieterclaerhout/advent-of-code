package day17

import (
	"fmt"
	"math"

	"github.com/pieterclaerhout/advent-of-code/day17/generics"
)

type Solver struct {
	Winds []int8
	Rocks []Rock
}

func NewSolver(winds []int8, rocks []Rock) Solver {
	return Solver{
		Winds: winds,
		Rocks: rocks,
	}
}

func (solver Solver) Solve(nrocks int64) int64 {
	world := make([]byte, 1000)

	var yOffset int64

	lru := generics.NewLruCache[State, Progress](100)

	prog := Progress{
		Y: 0,
		N: 0,
	}

	var windID int

	for ; prog.N < nrocks; prog.N++ {
		rockID := int(prog.N % int64(len(solver.Rocks)))

		x := 2
		y := int(prog.Y-yOffset) + 3

		for {
			w := int(solver.Winds[windID])
			windID = (windID + 1) % len(solver.Winds)
			if !solver.overlap(world, solver.Rocks[rockID], x+w, y) {
				x += w
			}
			// Drop
			if !solver.overlap(world, solver.Rocks[rockID], x, y-1) {
				y--
			} else {
				break
			}
		}

		var err error
		var droppedRows int
		world, droppedRows, err = solver.addRock(world, solver.Rocks[rockID], x, y)
		if err != nil {
			return 0
		}

		prog.Y = generics.Max(prog.Y, int64(y+solver.Rocks[rockID].Height)+yOffset)
		yOffset += int64(droppedRows)

		if int(prog.Y-yOffset) >= 4 {
			continue
		}

		s := solver.compressState(world, rockID, windID)

		cached, hit := lru.Get(s)
		if !hit {
			lru.Set(s, prog)
			continue
		}

		lru.Set(s, prog)

		deltaY := prog.Y - cached.Y
		deltaRock := prog.N - cached.N

		rocksLeft := nrocks - prog.N
		epochsLeft := (rocksLeft / deltaRock) - 1
		if epochsLeft < 1 {
			continue
		}

		prog.N += epochsLeft * deltaRock
		prog.Y += epochsLeft * deltaY
		yOffset += epochsLeft * deltaY
	}

	return prog.Y
}

func (solver Solver) overlap(world []byte, r Rock, X, Y int) bool {
	if Y < 0 || Y+r.Height > len(world) {
		return true
	}
	if X < 0 || X+r.Width > 7 {
		return true
	}

	for i, rockrow := range r.Shape {
		rockrow := rockrow >> X
		if world[Y+i]&rockrow != 0 {
			return true
		}
	}
	return false
}

func (solver Solver) addRock(world []byte, r Rock, x, y int) ([]byte, int, error) {
	minSize := y + r.Height + 10
	slack := len(world) - minSize
	if slack < 0 {
		world = append(world, make([]byte, len(world))...)
	}

	if err := solver.overwrite(world, r, x, y); err != nil {
		return nil, 0, err
	}

	var offset int
	for i := r.Height - 1; i >= 0; i-- {
		y := y + i
		const fullRow = math.MaxUint8 >> 1
		if world[y] == fullRow {
			offset = y + 1
			world = world[offset:]
			break
		}
	}

	return world, offset, nil
}

func (solver Solver) overwrite(world []byte, r Rock, X, Y int) error {
	if Y < 0 || Y+r.Height > len(world) {
		return fmt.Errorf("cannot write kernel into location i=%d. Want 0 <= i <= i+%d <= %d", Y, r.Height, len(world))
	}
	if X < 0 || X+r.Width > 7 {
		return fmt.Errorf("cannot write kernel into location j=%d. Want <= j <= j+%d <= 7", X, r.Width)
	}

	for i, rockrow := range r.Shape {
		world[Y+i] |= rockrow >> X
	}
	return nil
}

func (solver Solver) compressState(world []byte, rockID int, windID int) State {
	st := State{RockID: rockID, WindID: windID}
	copy(st.World[:], world)
	return st
}
