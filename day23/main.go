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

func SolvePartOne(ll DoublyLinkedList) {
	// seen := make(map[*Node]bool)
	p := ll.Head

	for range 100 {

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
			dst += 9
		}
		for cupVals[dst] {
			dst -= 1
			if dst <= 0 {
				dst += 9
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

	val := CollectValues(ll.Mp[1])
	fmt.Println(val)

}

func main() {

	// input := 389125467
	input := 614752839

	ll := CreateLinkedList(input)
	SolvePartOne(ll)
}
