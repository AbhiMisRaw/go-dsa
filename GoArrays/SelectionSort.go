package main

func SelectionSort(arr []int) {
	arrSize := len(arr)
	for i := 0; i < arrSize; i++ {
		smallest := i                      // let the smallest index is : i
		for j := i + 1; j < arrSize; j++ { // this loop will find the smallest element's index
			if arr[j] < arr[smallest] {
				smallest = j // if we find any smallest element then we assumed in previous, we will re-assign
			}
		}
		if arr[i] > arr[smallest] {
			arr[i], arr[smallest] = arr[smallest], arr[i]
		}

	}
}
