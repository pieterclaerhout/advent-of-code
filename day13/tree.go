package day13

type Tree struct {
	ValueLeaf int
	Elements  []*Tree
	Father    *Tree
}

func NewTree(parent *Tree) Tree {
	return Tree{
		ValueLeaf: -1,
		Elements:  []*Tree{},
		Father:    parent,
	}
}
