package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// generate head node
func New() *ListNode {
	return &ListNode{0, nil}
}

// insert data in 'i' position
func (head *ListNode) insert(data int, i int) bool {
	p := head
	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}
	if p == nil || j > i {
		fmt.Println("error")
		return false
	}
	s := &ListNode{Val: data}
	s.Next = p.Next
	p.Next = s
	return true
}

// 遍历
func (head *ListNode) traverse() {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func (head *ListNode) delete(i int) bool {
	p := head
	j := 1
	for p != nil && i > j {
		p = p.Next
		j++
	}
	if p == nil {
		return false
	}
	p.Next = p.Next.Next
	return true
}
// get 'i' position val
func (head *ListNode) get(i int) int {
	p := head.Next
	for j := 1; j <= i; j++ {
		if p == nil {
			return -1
		}else{
			p = p.Next
		}

	}
	return p.Val
}
// main
func main() {
	l := New()
	l.insert(1, 1)
	l.insert(2, 2)
	//l.insert(3, 5)
	//l.delete(2)
	fmt.Println(l.get(3))

	//l.traverse()
}
