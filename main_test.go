package main

import (
	"fmt"
	"hhping821/leetcode/list"
	"testing"
)

func Test_ListNode(t *testing.T) {
	data := []int{1, 3, 5, 6, 10, 28}

	l := new(list.ListNode)
	l.Init(data)
	l.Traverse()
	fmt.Println(l.Get(8))
	l.RemoveDataNode(1)
	l.Traverse()
	l.RemoveIndexNode(4)
	l.Traverse()
}
