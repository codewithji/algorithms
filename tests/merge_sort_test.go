package tests

import (
	"testing"

	algos "github.com/codewithji/go-algorithms/algorithms"
)

func TestMergeSort(t *testing.T) {
	sorted := algos.MergeSort([]int{9, 3, 5, 1, 0})
	answer := []int{0, 1, 3, 5, 9}
	for i := range sorted {
		if sorted[i] != answer[i] {
			t.Errorf("Test failed, expected %v, got %v", []int{0, 1, 3, 5, 9}, sorted)
		}
	}
}
