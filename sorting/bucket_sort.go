package sorting

var BucketSize = 5

func BucketSort(arr []int) {
	minValue, maxValue := arr[0], arr[0]
	for _, v := range arr {
		if v < minValue {
			minValue = v
		}
		if v > maxValue {
			maxValue = v
		}
	}

	buckets := make([][]int, (maxValue-minValue)/BucketSize+1)
	for _, v := range arr {
		index := int((v - minValue) / BucketSize)
		buckets[index] = append(buckets[index], v)
	}

	arr = arr[:0]
	//对每个桶内的数进行排序
	for _, bucket := range buckets {
		InsertionSort(bucket)
		//将排序好的数加入到arr中
		arr = append(arr, bucket...)
	}
}
