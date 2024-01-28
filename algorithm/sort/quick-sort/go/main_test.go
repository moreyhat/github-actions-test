package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	test_cases := [][][]int{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{4, 1, 3, 2}, []int{1, 2, 3, 4}},
		{[]int{4, 1, 3}, []int{1, 3, 4}},
		{[]int{1000, 2000, 3000, 4000, 5000}, []int{1000, 2000, 3000, 4000, 5000}},
		{[]int{-1, -3, -2, -4, -5}, []int{-5, -4, -3, -2, -1}},
		{[]int{-1, 2, -3, 4, -5}, []int{-5, -3, -1, 2, 4}},
		{[]int{-1, 2, -3, 4, -5}, []int{-5, -3, -1, 2, 4}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{1, 2, 3, 2, 1, 3, 2}, []int{1, 1, 2, 2, 2, 3, 3}},
		{[]int{5, 3, 9, 1, 6, 7, 2, 4, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{-100, 0, 100, -50, 50}, []int{-100, -50, 0, 50, 100}},
	}

	for _, test_case := range test_cases {
		input := test_case[0]
		expected := test_case[1]
		actual := QuickSort(input)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("The output from QuickSort is not expected. expected: %v, output: %v", expected, actual)
		}
	}
}
