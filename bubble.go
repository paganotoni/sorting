package sorting

func bubblesort(data []int) {
	nextIteration := true
	for nextIteration {
		nextIteration = false
		for i, value := range data {
			if i == 0 {
				continue
			}

			if data[i-1] > value {
				data[i] = data[i-1]
				data[i-1] = value

				nextIteration = true
				continue
			}
		}

		if nextIteration == false {
			break
		}
	}
}
