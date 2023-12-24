package main

func InsertionSort(arr []int) {
	arrSize := len(arr)
	for i := 1; i < arrSize; i++ {
		//var j int
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}

		}
	}
}
