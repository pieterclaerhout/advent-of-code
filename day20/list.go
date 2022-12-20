package day20

type Node struct {
	Value         int
	Moved         bool
	Left          *Node
	Right         *Node
	OriginalIndex int
}

func NewList(in []int) *Node {
	var front, prev *Node
	for i := 0; i < len(in); i++ {
		if i == 0 {
			front = &Node{Value: in[i], OriginalIndex: i}
			prev = front
			continue
		}

		n := &Node{Value: in[i], Left: prev, OriginalIndex: i}
		prev.Right = n
		prev = n
	}
	front.Left = prev
	prev.Right = front
	return front
}

func Mix(front *Node, count int) {
	cur := front
	ptr := front
	idx := 0
	for idx < count {
		tomove := cur.Value
		dir := 1
		if tomove < 0 {
			dir = -1
			tomove *= -1
		}

		tomove = tomove % (count - 1)
		// this also works.
		// for tomove/count != 0 {
		// 	tomove = tomove/count + tomove%count
		// }

		ptr = cur.Right
		cur.Moved = true

		for j := 0; j < tomove; j++ {
			tmpleft := cur.Left
			tmpright := cur.Right
			tmpright.Left = tmpleft
			tmpleft.Right = tmpright
			if dir == 1 {
				cur.Right = tmpright.Right
				cur.Left = tmpright
				tmpright.Right = cur
				cur.Right.Left = cur

			} else {
				cur.Left = tmpleft.Left
				cur.Right = tmpleft
				tmpleft.Left = cur
				cur.Left.Right = cur
			}
		}

		idx++
		for ptr.OriginalIndex != idx && idx < count {
			ptr = ptr.Right
		}
		cur = ptr
	}
}

func GroveCoordinates(front *Node) int {
	gc := []int{}

	cur := front
	for cur.Value != 0 {
		cur = cur.Right
	}

	for len(gc) < 3 {
		for i := 0; i < 1000; i++ {
			cur = cur.Right
		}
		gc = append(gc, cur.Value)
	}
	return gc[0] + gc[1] + gc[2]
}

func ApplyDecryptKey(front *Node, k int) {
	cur := front
	first := true
	for cur != front || first {
		first = false
		cur.Value *= k
		cur = cur.Right
	}
}
