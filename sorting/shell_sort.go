package sorting

func ShellSort(list []int) {
	for d := int(len(list) / 2); d > 0; d /= 2 {
		for i := d; i < len(list); i++ {
			for j := i; j >= d && list[j-d] > list[j]; j -= d {
				list[j], list[j-d] = list[j-d], list[j]
			}
		}
	}

}
