package sorting

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	cases := []struct {
		Data     []int
		Expected []int
	}{
		{[]int{5, 3, 7, 6, 2, 1, 4}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{2, 3, 4, 6, 5, 1, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{4, 3, 5}, []int{3, 4, 5}},
		{[]int{4, 3, -1}, []int{-1, 3, 4}},
		{[]int{4, 3, -1, -2}, []int{-2, -1, 3, 4}},
		{[]int{4, 3, -1, -2}, []int{-2, -1, 3, 4}},
	}

	for _, tcase := range cases {

		t.Run(fmt.Sprintf("%v", tcase.Data), func(t *testing.T) {
			target := make([]int, len(tcase.Data))
			copy(target, tcase.Data)

			bubblesort(target)
			if reflect.DeepEqual(target, tcase.Expected) == false {
				t.Errorf("Output for %v should be %v but instead was %v", tcase.Data, tcase.Expected, target)
			}
		})

	}
}

func BenchmarkBubble100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, 100)
		for j := 0; j < 100; j++ {
			data[j] = rand.Int()
		}

		bubblesort(data)
	}
}

func BenchmarkBubblesort(b *testing.B) {
	lengths := []int{10, 100, 1000, 10000}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("%v", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				data := make([]int, length)
				for j := 0; j < length; j++ {
					data[j] = rand.Int()
				}

				bubblesort(data)
			}
		})
	}
}
