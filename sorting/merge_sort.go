package sorting

func Merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}

func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	mid := len(list) / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return Merge(left, right)
}
