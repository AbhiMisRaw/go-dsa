package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func Create(data []int, head **Node) {
	var curr *Node = *head
	for i := 0; i < len(data); i++ {
		newNode := Node{Data: data[i], Next: nil}
		if curr == nil {
			*head = &newNode
			curr = &newNode
		} else {
			curr.Next = &newNode
			curr = curr.Next
		}
	}
}

func Display(head *Node) {
	var curr *Node = head
	var count int = 0
	for curr != nil {
		fmt.Printf("%d ", curr.Data)
		curr = curr.Next
		count++
	}
	if count == 0 {
		fmt.Println("No elements to Print")
	}
}

func DisplayRecursively(head *Node) {
	if head == nil {
		return
	} else {
		fmt.Print(" ", head.Data)
		DisplayRecursively(head.Next)
	}
}

func AddFirst(head **Node, data int) {
	newNode := Node{Data: data, Next: *head}
	*head = &newNode

}

func AddElements(head *Node) int {
	res := 0
	for p := head; p != nil; p = p.Next {
		res += p.Data
	}
	return res
}
