package main

func partition(arr []int, low int, high int) int {
	i := low

	for j := low; j < high; j++ {
		if arr[j] < arr[high] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		i := partition(arr, low, high)

		quickSort(arr, low, i-1)
		quickSort(arr, i+1, high)
	}
}
