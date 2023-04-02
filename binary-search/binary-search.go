package binary_search

type IBinarySearch interface {
	FindIndex([]int, int) int
}
type BinarySearch struct {
}

// iterative solution for sorted array
func (s *BinarySearch) FindIndex(arr []int, target int) int {
	mid, start, end := len(arr)/2, 0, len(arr)-1

	for start <= end {
		value := arr[mid]
		if value == target {
			return mid
		} else if value > target {
			end = mid - 1
			mid = (start + end) / 2
			continue
		}
		start = mid + 1
		mid = (start + end) / 2
	}
	return -1
}

func NewBinarySearch() IBinarySearch {
	return &BinarySearch{}
}
