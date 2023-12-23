package main

import "fmt"

func main() {
	arr := []int{12, 7, 18, 15, 24, 46, 5, 29, 33}
	fmt.Print("\nBefore Sorting :", arr)
	SelectionSort(arr)
	fmt.Print("\nAfter Sorting :", arr)
}
