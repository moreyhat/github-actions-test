package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{3, 5, 1, 6, 7}
	expected := []int{1, 3, 5, 6, 7}

	actual := QuickSort(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("The output from QuickSort is not expected. expected: %v, output: %v", expected, actual)
	}
}
