package day18

type Cube struct {
	X int
	Y int
	Z int
}

type Cubes []Cube

func (cubes Cubes) BuildMap() (map[Cube]bool, int) {
	cubesMap := map[Cube]bool{}
	area := 0

	for _, cube := range cubes {
		cubesMap[cube] = true
		area += 6

		if cubesMap[Cube{cube.X, cube.Y, cube.Z + 1}] {
			area -= 2
		}
		if cubesMap[Cube{cube.X, cube.Y, cube.Z - 1}] {
			area -= 2
		}
		if cubesMap[Cube{cube.X + 1, cube.Y, cube.Z}] {
			area -= 2
		}
		if cubesMap[Cube{cube.X - 1, cube.Y, cube.Z}] {
			area -= 2
		}
		if cubesMap[Cube{cube.X, cube.Y + 1, cube.Z}] {
			area -= 2
		}
		if cubesMap[Cube{cube.X, cube.Y - 1, cube.Z}] {
			area -= 2
		}
	}

	return cubesMap, area
}

func (cubes Cubes) BuildMap3D() CubesMap {
	cubesMap := [25][25][25]Node{}
	for _, c := range cubes {
		cubesMap[c.X+1][c.Y+1][c.Z+1].IsCube = true
	}
	return cubesMap
}
