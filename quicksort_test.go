package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {

	cases := []struct {
		Input  []int
		Output []int
	}{
		{[]int{5, 3, 7, 6, 2, 1, 4}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{2, 3, 4, 6, 5, 1, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{4, 3, 5}, []int{3, 4, 5}},
		{[]int{4, 3, -1}, []int{-1, 3, 4}},
		{[]int{4, 3, -1, -2}, []int{-2, -1, 3, 4}},
		{[]int{4, 3, -1, -2}, []int{-2, -1, 3, 4}},
	}

	for _, tcase := range cases {

		t.Run(fmt.Sprintf("%v", tcase.Input), func(t *testing.T) {
			result := quicksort(tcase.Input)
			if reflect.DeepEqual(result, tcase.Output) == false {
				t.Errorf("Output for %v should be %v but instead was %v", tcase.Input, tcase.Output, result)
			}
		})

	}

}
