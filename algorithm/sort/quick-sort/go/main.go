package main

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := int(len(arr) / 2)
	left := []int{}
	right := []int{}

	for i, item := range arr {
		if i != mid {
			if item > arr[mid] {
				right = append(right, item)
			} else {
				left = append(left, item)
			}
		}
	}

	sorted_left := QuickSort(left)
	sorted_right := QuickSort(right)

	return append(append(sorted_left, arr[mid]), sorted_right...)
}
