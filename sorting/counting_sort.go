package sorting

import "fmt"

func CountingSort(arr []int) {
	//参数处理
	if len(arr) == 0 {
		fmt.Printf("%s\n", "err: len(arr) == 0")
		return
	}

	//1.确定最大值
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	//由最大值构建计数桶
	buckets := make([]int, max+1)
	//将元素放到计数桶里面计数
	for _, v := range arr {
		buckets[v] += 1
	}
	//遍历计数桶取出元素排序
	arrIndex := 0
	for i, v := range buckets {
		for v > 0 && arrIndex < len(arr) {
			arr[arrIndex] = i
			arrIndex++
			v--
		}
	}
}
