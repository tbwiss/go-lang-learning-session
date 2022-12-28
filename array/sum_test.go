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

func TestSumAllTails(t *testing.T) {
	CheckSums := func(sum []int, want []int, t testing.TB) {
		t.Helper()
		if !reflect.DeepEqual(sum, want) {
			t.Errorf("sum %v, want %v", sum, want)
		}
	}

	t.Run("sum all tails", func(t *testing.T) {
		sum := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		CheckSums(sum, want, t)
	})

	t.Run("sum all tails with empty slice", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{3, 4, 6})
		want := []int{0, 10}

		CheckSums(sum, want, t)
	})
}
