package main

import "fmt"

func main() {
	arr := []int{12, 7, 18, 15, 49, 31, 11, 30, 29, 33}
	fmt.Print("\nBefore Sorting :", arr)
	InsertionSort(arr)
	fmt.Print("\nAfter Sorting :", arr)
}
