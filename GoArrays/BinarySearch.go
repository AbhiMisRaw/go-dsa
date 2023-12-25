package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	for start < end {
		printArray(arr, start, end)
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	return -1
}
func printArray(arr []int, start int, end int) {
	for i := start; i <= end; i++ {
		fmt.Print(" ", arr[i])
	}
	fmt.Println()
}
