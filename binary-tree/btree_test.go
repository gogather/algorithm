package binary_tree

import (
	"fmt"
	"testing"
)

type Integer int

func (i Integer) Compare(value interface{}) int {
	v := int(value.(Integer))
	if v < int(i) {
		return 1
	} else if v == int(i) {
		return 0
	} else {
		return -1
	}
}

func TestNewBTree(t *testing.T) {
	bt := NewBTree()
	bt.Insert(Integer(4))
	bt.Insert(Integer(41))
	bt.Insert(Integer(7))
	bt.Insert(Integer(13))
	bt.Insert(Integer(3))
	bt.Insert(Integer(1))
	bt.Insert(Integer(42))

	fmt.Println(bt)
}
