package quicksort

func quicksort(data []int) []int {

	if len(data) < 2 {
		return data
	}

	pivotIdx := len(data) - 1
	pivot := data[pivotIdx]
	cutIndex := -1

	j := pivotIdx
	i := 0

	for true {

		iData := data[i]
		jData := data[j]

		needsSwitch := data[i] > data[j]

		if needsSwitch {
			tmp := data[j]
			data[j] = data[i]
			data[i] = tmp
		}

		if jData > pivot {
			j--
		}

		if iData < pivot {
			i++
		}

		if i == j {
			cutIndex = i
			quicksort(data[0:cutIndex])
			quicksort(data[cutIndex:len(data)])
			break
		}

	}

	return data
}
