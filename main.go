package main

import (
	"fmt"
	"hhping821/leetcode/list"
)

func main() {
	data := []int{1, 3, 5, 6, 10, 28}

	l := new(list.ListNode)
	l.Init(data)
	l.Traverse()
	fmt.Println(l.Get(5))
}
