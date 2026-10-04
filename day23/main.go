package main

import (
	"fmt"
)

// Double Linked List Node
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Mp   map[int]*Node
	Head *Node
}

func (ll *DoublyLinkedList) AddRemainingNodes() {
	tail := ll.Head.Prev
	for val := 10; val <= 1e6; val++ {
		ll.Mp[val] = &Node{
			Val:  val,
			Next: nil,
			Prev: nil,
		}

		tail.Next, ll.Mp[val].Prev = ll.Mp[val], tail

		tail = tail.Next
	}

	ll.Head.Prev = tail
	tail.Next = ll.Head

}

func CreateLinkedList(input int) DoublyLinkedList {
	var nxt *Node
	var tail *Node
	mp := make(map[int]*Node)
	// Building the DLL tail --> head
	for input > 0 {
		digit := input % 10
		input = input / 10

		mp[digit] = new(Node{
			Val:  digit,
			Next: nxt,
			Prev: nil,
		})

		if tail == nil {
			tail = mp[digit]
		}

		if mp[digit].Next != nil {
			mp[digit].Next.Prev = mp[digit]
		}

		nxt = mp[digit]
	}

	nxt.Prev, tail.Next = tail, nxt

	return DoublyLinkedList{
		Mp:   mp,
		Head: nxt,
	}

}

func CollectValues(head *Node) int {
	var res int
	seen := make(map[*Node]bool)
	p := head

	// skip 1 node
	seen[p] = true
	p = p.Next

	for !seen[p] {
		res = res*10 + p.Val
		seen[p] = true
		p = p.Next
	}

	return res
}

func Move(ll *DoublyLinkedList, turns int, maxVal int) {
	p := ll.Head

	for range turns {

		// Take next 3 cups
		l := p.Next
		cupVals := map[int]bool{}
		r := p
		for range 3 {
			r = r.Next
			cupVals[r.Val] = true
		}

		// Determine destination cup
		dst := p.Val - 1
		if dst == 0 {
			dst += maxVal
		}
		for cupVals[dst] {
			dst -= 1
			if dst <= 0 {
				dst += maxVal
			}
		}

		// fmt.Println(l.Val, r.Val, cupVals, dst, ll.Mp[dst])

		// Unhook the 3 cups
		l.Prev.Next, r.Next.Prev = r.Next, l.Prev

		// Rehook the 3 cups
		dstNode := ll.Mp[dst]
		dstNodeNxt := dstNode.Next

		dstNode.Next, l.Prev = l, dstNode
		dstNodeNxt.Prev, r.Next = r, dstNodeNxt

		p = p.Next
	}

}

func main() {

	// input := 389125467
	input := 614752839

	// Part One
	ll := CreateLinkedList(input)
	Move(&ll, 100, 9)
	res := CollectValues(ll.Mp[1])
	fmt.Println(res)

	ll2 := CreateLinkedList(input)
	ll2.AddRemainingNodes()
	Move(&ll2, 10e6, 1e6)
	res2 := ll2.Mp[1].Next.Val * ll2.Mp[1].Next.Next.Val
	fmt.Println(res2)

}
