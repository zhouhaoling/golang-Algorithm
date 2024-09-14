package sorting

func QuickSort(list []int, low, high int) {
	if high > low {
		loc := partition(list, low, high)
		//对左边进行快排
		QuickSort(list, low, loc-1)
		//对右边进行快排
		QuickSort(list, loc+1, high)
	}
}

func partition(list []int, begin, end int) int {
	i := begin + 1 //将list[begin]作为pivot key
	j := end
	for i < j {
		if list[i] > list[begin] {
			list[i], list[j] = list[j], list[i] //大于基准数，与数组最后一个交换,下标向左移动
			j--
		} else {
			i++ // 小于或等于基准数，下标向左移动
		}
	}
	//for循环结束后，i=j
	if list[i] >= list[begin] {
		i--
	}
	list[begin], list[i] = list[i], list[begin] //将pivot key放到合适的位置
	return i
}
