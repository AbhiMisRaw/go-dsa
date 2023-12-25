package main

import "fmt"

func main() {
	arr := []int{12, 7, 18, 15, 49, 31, 11, 30, 29, 33}
	fmt.Print("\nBefore Sorting :", arr)
	InsertionSort(arr)
	fmt.Print("\nAfter Sorting :", arr)
	fmt.Println("Searching element: ")
	element := 18
	index := BinarySearch(arr, element)
	fmt.Println(" element ", element, "is present in index : ", index)

}
