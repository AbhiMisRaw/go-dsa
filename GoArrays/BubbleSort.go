package main

func BubbleSort(arr []int) {
	var size int = len(arr)
	for i := 0; i < size; i++ {
		for j := 1; j < size; j++ {
			if arr[j] < arr[j-1] {
				//swap(arr, j, j-1)
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
