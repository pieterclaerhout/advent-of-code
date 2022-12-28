package day09

type Point struct {
	X int
	Y int
}

func (tail Point) AdjustTail1(head Point) Point {
	newTail := tail

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{-2, 1},
		Point{-1, 2},
		Point{0, 2},
		Point{1, 2},
		Point{2, 1}:
		newTail.Y++
	}

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{1, 2},
		Point{2, 1},
		Point{2, 0},
		Point{2, -1},
		Point{1, -2}:
		newTail.X++
	}

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{2, -1},
		Point{1, -2},
		Point{0, -2},
		Point{-1, -2},
		Point{-2, -1}:
		newTail.Y--
	}

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{-1, -2},
		Point{-2, -1},
		Point{-2, -0},
		Point{-2, 1},
		Point{-1, 2}:
		newTail.X--
	}

	return newTail
}

func (tail Point) AdjustTail2(head Point) Point {
	newTail := tail

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{-2, 1},
		Point{-1, 2},
		Point{0, 2},
		Point{1, 2},
		Point{2, 1},
		Point{2, 2},
		Point{-2, 2}:
		newTail.Y++
	}

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{1, 2},
		Point{2, 1},
		Point{2, 0},
		Point{2, -1},
		Point{1, -2},
		Point{2, 2},
		Point{2, -2}:
		newTail.X++
	}

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{-2, -2},
		Point{2, -1},
		Point{1, -2},
		Point{0, -2},
		Point{-1, -2},
		Point{-2, -1},
		Point{2, -2}:
		newTail.Y--
	}

	switch (Point{head.X - tail.X, head.Y - tail.Y}) {
	case
		Point{-2, -2},
		Point{-1, -2},
		Point{-2, -1},
		Point{-2, -0},
		Point{-2, 1},
		Point{-1, 2},
		Point{-2, 2}:
		newTail.X--
	}

	return newTail
}
