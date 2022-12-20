package day15

type Sensor struct {
	Position Point
	Beacon   Point
}

type Sensors []Sensor

func (sensors Sensors) Coverage(y int) int {
	coverage := map[int]bool{}

	for _, sensor := range sensors {
		dist := sensors.manhattanDistance(sensor.Position, sensor.Beacon)
		yDist := sensors.abs(y - sensor.Position.Y)
		xDist := dist - yDist
		for x := sensor.Position.X - xDist; x <= sensor.Position.X+xDist; x++ {
			if sensor.Beacon.Y != y || x != sensor.Beacon.X {
				coverage[x] = true
			}
		}
	}

	return len(coverage)
}

func (sensors Sensors) Coverage2D(side int) int {
	for _, sensor := range sensors {
		dist := sensors.manhattanDistance(sensor.Position, sensor.Beacon) + 1

		p := Point{
			X: sensor.Position.X,
			Y: sensor.Position.Y - dist,
		}

		// down - right
		for i := 0; i < dist; i++ {
			if p.X < 0 || p.X > side || p.Y < 0 || p.Y > side {
				continue
			}
			if !sensors.isInsideAny(p.X, p.Y) {
				return p.X*4000000 + p.Y
			}
			p.X++
			p.Y++
		}
		// down - left
		for i := 0; i < dist; i++ {
			if p.X < 0 || p.X > side || p.Y < 0 || p.Y > side {
				continue
			}
			if !sensors.isInsideAny(p.X, p.Y) {
				return p.X*4000000 + p.Y
			}
			p.X--
			p.Y++
		}
		// up - left
		for i := 0; i < dist; i++ {
			if p.X < 0 || p.X > side || p.Y < 0 || p.Y > side {
				continue
			}
			if !sensors.isInsideAny(p.X, p.Y) {
				return p.X*4000000 + p.Y
			}
			p.X--
			p.Y--
		}
		// up - right
		for i := 0; i < dist; i++ {
			if p.X < 0 || p.X > side || p.Y < 0 || p.Y > side {
				continue
			}
			if !sensors.isInsideAny(p.X, p.Y) {
				return p.X*4000000 + p.Y
			}
			p.X++
			p.Y--
		}
	}

	return 0
}

func (sensors Sensors) manhattanDistance(a Point, b Point) int {
	return sensors.abs(a.X-b.X) + sensors.abs(a.Y-b.Y)
}

func (sensors Sensors) isInsideAny(x int, y int) bool {
	for _, sensor := range sensors {
		if sensors.manhattanDistance(Point{X: x, Y: y}, sensor.Position) <= sensors.manhattanDistance(sensor.Position, sensor.Beacon) {
			return true
		}
	}
	return false
}

func (sensors Sensors) abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
