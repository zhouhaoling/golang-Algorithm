package main

import "fmt"

// 二分查找函数
func binarySearch(arr []int, target int) (int, int) {
	low, high := 0, len(arr)-1
	number := 0
	for low <= high {
		number++
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid, number // 找到目标值，返回索引
		} else if arr[mid] < target {
			low = mid + 1 // 目标值在右半部分，缩小搜索范围
		} else {
			high = mid - 1 // 目标值在左半部分，缩小搜索范围
		}
	}

	return -1, number // 目标值不存在
}

func main() {
	// 示例数组，必须是已排序的数组
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

	// 要查找的目标值
	target := 100

	// 调用二分查找算法
	index, n := binarySearch(arr, target)
	fmt.Println("运行的次数：", n)
	if index != -1 {
		fmt.Printf("目标值 %d 在数组中的索引是 %d\n", target, index)
	} else {
		fmt.Printf("目标值 %d 不存在于数组中\n", target)
	}
}
