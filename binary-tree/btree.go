package binary_tree

import "fmt"

// comparable node value
type Comparable interface {
	Compare(value Comparable) int
}

// binary tree node
type BTreeNode struct {
	LNode *BTreeNode
	RNode *BTreeNode
	Value Comparable
}

func (btn *BTreeNode) String() string {
	return fmt.Sprintf("Node: [%v]", btn.Value)
}

// binary tree
type BTree struct {
	TopNode    *BTreeNode
	NodeNumber int
}

func (bt *BTree) String() string {
	return fmt.Sprintf("Tree Top Node: [%v], Total Node Number: %d", bt.TopNode.Value, bt.NodeNumber)
}

func (bt *BTree) Insert(value Comparable) {
	node := &BTreeNode{
		LNode: nil,
		RNode: nil,
		Value: value,
	}

	// search first
	edgeNode := bt.SearchEdge(value)

	if edgeNode == nil {
		bt.TopNode = node
		bt.NodeNumber++
	} else {
		// then, insert
		compare := edgeNode.Value.Compare(value)
		if compare == 0 {
			// nop
		} else if compare > 0 {
			// put left
			edgeNode.LNode = node
			bt.NodeNumber++
		} else if compare < 0 {
			// put right
			edgeNode.RNode = node
			bt.NodeNumber++
		}
	}
}

func (bt *BTree) Search(value Comparable) (node *BTreeNode) {
	edge := bt.SearchEdge(value)
	if edge != nil && edge.Value.Compare(value) == 0 {
		return edge
	} else {
		return nil
	}
}

func (bt *BTree) SearchEdge(value Comparable) (edge *BTreeNode) {
	if bt.TopNode == nil {
		return nil
	}

	currentNode := bt.TopNode
	for {
		compare := currentNode.Value.Compare(value)
		if compare == 0 {
			return currentNode
		} else if compare > 0 {
			// get left
			if currentNode.LNode == nil {
				return currentNode
			} else {
				currentNode = currentNode.LNode
			}
		} else if compare < 0 {
			// get right
			if currentNode.RNode == nil {
				return currentNode
			} else {
				currentNode = currentNode.RNode
			}
		}
	}
}

func NewBTree() (bt *BTree) {
	return &BTree{}
}
