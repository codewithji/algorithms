package algorithms

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2

	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func merge(left, right []int) []int {
	size := len(left) + len(right)
	var i, j int
	slice := make([]int, 0, size)

	for i < len(left) || j < len(right) {
		if i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				slice = append(slice, left[i])
				i++
			} else {
				slice = append(slice, right[j])
				j++
			}
		} else {
			if i < len(left) {
				slice = append(slice, left[i])
				i++
			} else {
				slice = append(slice, right[j])
				j++
			}
		}
	}

	return slice
}
