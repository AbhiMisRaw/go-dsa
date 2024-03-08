package main

import "fmt"

func main() {
	var head *Node
	arr := []int{1, 2, 3, 4, 5, 6}
	Create(arr, &head)
	fmt.Println("Linked List is created")
	fmt.Println("Content of Linked List : ")
	fmt.Println("Address of HEAD", &head)
	AddFirst(&head, 10)
	AddFirst(&head, 20)
	AddFirst(&head, 30)
	Display(head)
	fmt.Println("Printed list using Recursive display function")
	DisplayRecursively(head)
	fmt.Println("Hello world")
	sum := AddElements(head)
	fmt.Println(sum)
}
