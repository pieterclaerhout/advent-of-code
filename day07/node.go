package day07

type Node struct {
	Name     string
	Size     int
	Children map[string]*Node
	Parent   *Node
}

func NewNode(name string, size int, parent *Node) Node {
	return Node{
		Name:     name,
		Size:     size,
		Children: map[string]*Node{},
		Parent:   parent,
	}
}

func NewDir(name string, parent *Node) Node {
	return NewNode(name, 0, parent)
}

func NewFile(name string, size int, parent *Node) Node {
	return NewNode(name, size, parent)
}

func (node *Node) SumLessThan100000() int {
	sum := 0
	for _, child := range node.Children {
		if len(child.Children) == 0 {
			continue
		} else {
			sum += child.SumLessThan100000()
			if child.Size <= 100000 {
				sum += child.Size
			}
		}
	}
	return sum
}

func (node *Node) ComputeSize() int {
	for _, child := range node.Children {
		if len(child.Children) == 0 {
			node.Size += child.Size
		} else {
			node.Size += child.ComputeSize()
		}
	}
	return node.Size
}

func (node *Node) FindSmallerDir(neededSpace int) int {
	minSpace := 70000000
	for _, child := range node.Children {
		if len(child.Children) == 0 {
			continue
		} else {
			if child.Size < minSpace && child.Size >= neededSpace {
				minSpace = child.Size
			}
			subDirSize := child.FindSmallerDir(neededSpace)
			if subDirSize < minSpace && subDirSize >= neededSpace {
				minSpace = subDirSize
			}
		}
	}
	return minSpace
}
