package sorting

// InsertionSort 插入排序
func InsertionSort(list []int) {
	var i, j int
	//从1开始，因为0已经是有序的
	for i = 1; i < len(list); i++ {
		//有序中最小的和当前元素比较，如果比有序中最小的小，则插入有序区
		for j = 0; j < i; j++ {
			if list[j] > list[i] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

/*
*demo: 4 2 9 1
*first run: j = 0, list[j] = 4, i = 1, list[i] = 2; ——》  2 4
*second run: ——》 2 4 9
*third run: ——》 1 4 9 2 —— 1 2 9 4 —— 1 2 4 9
 */
