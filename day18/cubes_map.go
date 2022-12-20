package day18

type CubesMap [25][25][25]Node

func (cubesMap *CubesMap) DfsSearch(c Cube) {
	cubesMap[c.X][c.Y][c.Z].Visited = true
	L := len(cubesMap)

	for _, p := range []Cube{
		{
			X: c.X,
			Y: c.Y,
			Z: c.Z + 1,
		},
		{
			X: c.X,
			Y: c.Y,
			Z: c.Z - 1,
		},
		{
			X: c.X,
			Y: c.Y + 1,
			Z: c.Z,
		},
		{
			X: c.X,
			Y: c.Y - 1,
			Z: c.Z,
		},
		{
			X: c.X + 1,
			Y: c.Y,
			Z: c.Z,
		},
		{
			X: c.X - 1,
			Y: c.Y,
			Z: c.Z,
		},
	} {
		if p.X < 0 || p.X >= L || p.Y < 0 || p.Y >= L || p.Z < 0 || p.Z >= L {
			continue
		}

		if !cubesMap[p.X][p.Y][p.Z].IsCube && !cubesMap[p.X][p.Y][p.Z].Visited {
			cubesMap.DfsSearch(p)
		}
	}
}

func (cubesMap *CubesMap) Area() int {
	area := 0
	L := len(cubesMap)

	for x := range cubesMap {
		for y := range cubesMap[x] {
			for z := range cubesMap[x][y] {
				if !cubesMap[x][y][z].IsCube {
					continue
				}

				for _, p := range []Cube{
					{
						X: x,
						Y: y,
						Z: z + 1,
					},
					{
						X: x,
						Y: y,
						Z: z - 1,
					},
					{
						X: x,
						Y: y + 1,
						Z: z,
					},
					{
						X: x,
						Y: y - 1,
						Z: z,
					},
					{
						X: x + 1,
						Y: y,
						Z: z,
					},
					{
						X: x - 1,
						Y: y,
						Z: z,
					},
				} {
					if p.X < 0 || p.X >= L || p.Y < 0 || p.Y >= L || p.Z < 0 || p.Z >= L {
						continue
					}
					if cubesMap[p.X][p.Y][p.Z].Visited {
						area += 1
					}
				}
			}
		}
	}
	return area
}
