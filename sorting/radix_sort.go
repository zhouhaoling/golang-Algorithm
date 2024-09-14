package sorting

import "math"

func RadixSort(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	//获取数组的最大元素，用于判断最大位数
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	digitNum := digit(maxVal)
	//按数字位数排序
	for i := 0; i < digitNum; i++ {
		buckets := make([][]int, 10)
		for j := 0; j < n; j++ {
			digitVal := int(math.Pow10(i))
			index := nums[j] / digitVal % 10
			buckets[index] = append(buckets[index], nums[j])
		}
		nums = []int{}
		for j := 0; j < 10; j++ {
			nums = append(nums, buckets[j]...)
		}
	}
	return nums
}

/* 获取元素的位数 */
func digit(maxVal int) int {
	digitNum := 0
	for maxVal > 0 {
		digitNum++
		maxVal /= 10
	}
	return digitNum
}
