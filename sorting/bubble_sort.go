package sorting

func BubbleSort(list []int) {
	n := len(list)

	//n-1轮迭代
	for i := n - 1; i > 0; i-- {
		didSwap := false
		//从第一位开始迭代
		for j := 0; j < i; j++ {
			//前面的数比后面大，交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		//如果没有交换，说明已经排好序了
		if !didSwap {
			return
		}
	}
}
