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
	p := l
	if p == nil {
		return false
	}
	i := 1

	for p != nil && i < pos {
		p = p.Next
		i++
	}

	if p == nil || i > pos {
		return false
	}
	n := &ListNode{Data: data}
	n.Next = p.Next
	p.Next = n
	return true
}

//按索引值删除节点
func (l *ListNode) RemoveIndexNode(i int) bool {
	p := l
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		return false
	}
	p.Next = p.Next.Next
	return true
}

//按数据删除节点
func (l *ListNode) RemoveDataNode(data int) bool {
	p := l
	for nil != p {
		if p.Next != nil && p.Next.Data == data {
			p.Next = p.Next.Next
			return true
		}
		p = p.Next
	}
	return false
}

//打印链表
func (l *ListNode) Traverse() {
	p := l.Next
	for nil != p {
		fmt.Println(p.Data)
		p = p.Next
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
	p := l.Next
	for j := 1; j < i; j++ {
		if nil == p {
			//表示返回错误
			return 0, errors.New("未找到节点值")
		}
		p = p.Next
	}

	return p.Data, nil
}
