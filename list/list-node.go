package list

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Data int
	Next *ListNode
}

//添加一个节点，指定位置，默认添到在头部位置
func (l *ListNode) Insert(pos, data int) bool {
	if l == nil {
		return false
	}
	i := 1

	for l != nil && i < pos {
		l = l.Next
		i++
	}

	if l == nil || i > pos {
		return false
	}
	n := &ListNode{Data: data}
	n.Next = l.Next
	l.Next = n
	return true
}

//按索引值删除节点
func (l *ListNode) RemoveIndexNode(i int) bool {
	j := 1
	for nil != l && j < i {
		l = l.Next
		j++
	}
	if nil == l || j > i {
		return false
	}
	l.Next = l.Next.Next
	return true
}

//按数据删除节点
func (l *ListNode) RemoveDataNode(data int) bool {
	for nil != l {
		if l.Next != nil && l.Next.Data == data {
			l.Next = l.Next.Next
			return true
		}
		l = l.Next
	}
	return false
}

//打印链表
func (l *ListNode) Traverse() {
	if l == nil {
		return
	}
	for nil != l {
		fmt.Println(l.Data)
		l = l.Next
	}
	fmt.Println("--------done----------")
}

func (l *ListNode) Init(data []int) {
	for k, v := range data {
		l.Insert(k+1, v)
	}
}

// 获取链表中的第i个元素，复杂度为o(n)
func (l *ListNode) Get(i int) (int, error) {
	for j := 1; j < i; j++ {
		if nil == l {
			//表示返回错误
			return 0, errors.New("未找到节点值")
		}
		l = l.Next
	}

	return l.Data, nil
}

//删除有序链表中的重复值
func (l *ListNode) DeleteDuplicates() {
	for l != nil && l.Next != nil {
		if l.Data == l.Next.Data {
			l.Next = l.Next.Next
		} else {
			l = l.Next
		}
	}
}

//反正链表数据
func (l *ListNode) ReversalList() *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	point := new(ListNode)
	for l != nil {
		node := l
		l = l.Next
		node.Next = point.Next
		point.Next = node
	}
	return point
}
