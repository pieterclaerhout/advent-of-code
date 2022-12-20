package day17

type Input struct {
	rocks    [PATTERN_SIZE][7]bool
	time     int
	numRocks int
}

type Cached struct {
	numRocks int
	height   int
}

type Coords struct {
	x int
	y int
}

const PATTERN_SIZE = 20

type Simulation struct {
}

func (simulation Simulation) runSimulation(input string, maxRocks int) int {
	rocks := [][7]bool{}
	height := -1
	numRocks := 0
	time := 0
	for numRocks < maxRocks {
		rock := simulation.spawnRock(numRocks, height+1)
		numRocks++

		rock = simulation.pushRock(input, time, rocks, rock)
		time++

		for simulation.canFall(rocks, rock) {
			rock = simulation.fall(rock)
			rock = simulation.pushRock(input, time, rocks, rock)
			time++
		}

		height = simulation.stopRock(&rocks, rock, height)
	}
	return height
}

func (simulation Simulation) runSimulation2(input string, maxRocks int) int {
	rocks := [][7]bool{}
	height := -1
	numRocks := 0
	time := 0
	cache := map[Input]Cached{}
	offset := 0
	for numRocks < maxRocks {
		rock := simulation.spawnRock(numRocks, height+1)
		numRocks++
		rock = simulation.pushRock(input, time, rocks, rock)
		time++

		for simulation.canFall(rocks, rock) {
			rock = simulation.fall(rock)
			rock = simulation.pushRock(input, time, rocks, rock)
			time++
		}
		height = simulation.stopRock(&rocks, rock, height)

		if cache != nil && len(rocks) > PATTERN_SIZE {
			key := Input{
				*(*[20][7]bool)(rocks[len(rocks)-20:]),
				time % len(input),
				numRocks % 5,
			}

			if cached, ok := cache[key]; ok {
				toGo := maxRocks - numRocks
				repetitions := toGo / (numRocks - cached.numRocks)
				offset = repetitions * (height - cached.height)
				numRocks += repetitions * (numRocks - cached.numRocks)
				cache = nil
			}
			if cache != nil {
				cache[key] = Cached{numRocks, height}
			}
		}
	}

	return height + offset
}

func (simulation Simulation) spawnRock(rockNum, height int) []Coords {
	switch rockNum % 5 {
	case 0:
		return []Coords{{2, height + 3}, {3, height + 3}, {4, height + 3}, {5, height + 3}}
	case 1:
		return []Coords{{3, height + 3}, {2, height + 4}, {3, height + 4}, {4, height + 4}, {3, height + 5}}
	case 2:
		return []Coords{{2, height + 3}, {3, height + 3}, {4, height + 3}, {4, height + 4}, {4, height + 5}}
	case 3:
		return []Coords{{2, height + 3}, {2, height + 4}, {2, height + 5}, {2, height + 6}}
	case 4:
		return []Coords{{2, height + 3}, {3, height + 3}, {2, height + 4}, {3, height + 4}}
	}
	return []Coords{}
}

func (simulation Simulation) pushRock(input string, i int, rocks [][7]bool, rock []Coords) []Coords {
	xChange := 1
	if simulation.getJet(input, i) == '<' {
		xChange = -1
	}
	newRock := []Coords{}
	for _, r := range rock {
		newCoords := Coords{r.x + xChange, r.y}
		if !simulation.isEmpty(rocks, newCoords) {
			return rock
		} else {
			newRock = append(newRock, newCoords)
		}
	}
	return newRock
}

func (simulation Simulation) getJet(input string, time int) byte {
	return input[time%len(input)]
}

func (simulation Simulation) isEmpty(rocks [][7]bool, c Coords) bool {
	if c.x >= 7 || c.x < 0 || c.y < 0 {
		return false
	}
	if c.y >= len(rocks) {
		return true
	}
	return !rocks[c.y][c.x]
}

func (simulation Simulation) canFall(rocks [][7]bool, rock []Coords) bool {
	if rock[0].y < 0 {
		return false
	}
	for _, r := range rock {
		if !simulation.isEmpty(rocks, Coords{r.x, r.y - 1}) {
			return false
		}
	}
	return true
}

func (simulation Simulation) fall(rock []Coords) []Coords {
	newRock := []Coords{}
	for _, r := range rock {
		newRock = append(newRock, Coords{r.x, r.y - 1})
	}
	return newRock
}

func (simulation Simulation) stopRock(rocks *[][7]bool, rock []Coords, height int) int {
	for _, r := range rock {
		if r.y > height {
			height = r.y
		}
		for r.y >= len(*rocks) {
			(*rocks) = append((*rocks), [7]bool{false, false, false, false, false, false, false})
		}
		(*rocks)[r.y][r.x] = true
	}
	return height
}
