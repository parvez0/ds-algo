package cord

type Node interface {
	Lenght() int
	HasChildren() bool
}

type InternalNode struct {
	Len   int
	Left  Node
	Right Node
}

func(i InternalNode) Lenght() int {
	return i.Len
}

func(i InternalNode) HasChildren() bool {
	return true
}

type LeafNode struct {
	Len   int
	Value []rune 
}

func(l LeafNode) Lenght() int {
	return l.Len
}

func(l LeafNode) HasChildren() bool {
	return false
}

type Tree struct {
	Root Node
}

func(t *Tree) New(root Node) {
	t.Root = root
}

func(t *Tree) FindIndex(i int) (bool, rune) {
	if t.Root.Lenght() < i {
		return false, ' '
	}
	cursor := t.Root
	for cursor != nil {
		if !cursor.HasChildren() {
			return true, cursor.(LeafNode).Value[i-1]
		}
		internalNode := cursor.(InternalNode)
		if internalNode.Left.Lenght() >= i {
			cursor = internalNode.Left
		} else {
			cursor, i = internalNode.Right, i-internalNode.Left.Lenght()
		}
	}
	return false, ' '
}