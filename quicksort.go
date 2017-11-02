package sorting

func quicksort(data []int) []int {

	if len(data) < 2 {
		return data
	}

	pivot := data[len(data)-1]
	cutIndex := -1

	j := len(data) - 1
	i := 0

	for true {

		ivalue := data[i]
		jvalue := data[j]

		needsSwitch := data[i] > data[j]

		if needsSwitch {
			tmp := data[j]
			data[j] = data[i]
			data[i] = tmp
		}

		if jvalue > pivot {
			j--
		}

		if ivalue < pivot {
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
