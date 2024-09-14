package sorting

func SelectionSort(list []int) {
	n := len(list)
	//n-1轮迭代
	for i := 0; i < n-1; i++ {
		minValue := list[i] //最小数
		minIndex := i       //最小数下标
		for j := i + 1; j < n; j++ {
			if list[j] < minValue {
				minValue = list[j]
				minIndex = j
			}
		}
		//一轮循环结束，将最小数放到i位置
		if i != minIndex {
			list[minIndex] = list[i]
			list[i] = minValue
		}
	}
}
