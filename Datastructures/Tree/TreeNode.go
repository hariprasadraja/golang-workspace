package main

// TreeType defines the Type of Tree Data structure.
type TreeType string

// Available Tree Types
const (
	BinaryTree TreeType = "BinaryTree"
)

// TreeNode defines a Node in Tree data structure.
// It is common for all type of tree implementation.
type treeNode struct {
	Parent       *treeNode
	LeftChild    []*treeNode
	RightChild   []*treeNode
	LeftSibling  []*treeNode
	RightSibling []*treeNode
	IsLeafNode   bool
	Degree       uint
	Level        uint
	Height       uint
	Depth        uint
	Message      interface{}
	Path         []string
}

// Tree represents a tree data structure.
type Tree struct {
	root    *treeNode
	maxNode uint
	Height  uint
	Depth   uint
	Type    TreeType
}

//BTree is said to be binary tree if it has maximum of two nodes.
func BTree() (bt Tree) {
	bt = Tree{
		maxNode: 2,
		Type:    BinaryTree,
	}

	return
}

// Insert msg in Tree data structure.
func (t *Tree) Insert(msg interface{}) {
	if t.root != nil {

		return
	}

	t.root = &treeNode{
		Message: msg,
	}
}

func (node *treeNode) insert(msg interface{}) {
}

func main() {
	binaryTree := BTree()
	binaryTree.

}
