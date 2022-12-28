package array

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("sum as array, array size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		want := 15

		if sum != want {
			t.Errorf("sum %d, want %d, given %v", sum, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum as array, array size 5", func(t *testing.T) {
		sum := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(sum, want) {
			t.Errorf("sum %v, want %v", sum, want)
		}
	})
}
