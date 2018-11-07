package list

import (
	"fmt"
	"testing"
)

func Test_ListNode(t *testing.T) {
	data := []int{1, 3, 3, 5, 5, 6, 10, 10, 10, 10, 10, 11, 28}

	l := new(ListNode)
	l.Init(data)
	l.Traverse()
	fmt.Println(l.Get(2))
	l.RemoveDataNode(4)
	l.Traverse()
	l.DeleteDuplicates()
	l.Traverse()
	l = l.ReversalList()
	l.Traverse()
}
