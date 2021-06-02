package main

import "fmt"

type Node struct{
	next *Node
	val int
}
func isLooped(h *Node) bool {
	t := h
	r := h

	for r.next != nil {
		t = t.next
		if r.next.next != nil {
			r = r.next.next
		}
		if t == r {
			return true
		}
	}
	return false
}


func main() {
	l5 := Node{nil, 5}
	l4 := Node{&l5, 4}
	l3 := Node{&l4, 3}
	l2 := Node{&l3, 2}
	l1 := Node{&l2, 1}

	l5.next = &l3

	fmt.Println(isLooped(&l1))

}
