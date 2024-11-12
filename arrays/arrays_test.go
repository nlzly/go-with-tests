package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, expected, received []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, received) {
			t.Errorf("expected %v received %v", expected, received)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		received := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, expected, received)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		received := SumAllTails([]int{}, []int{0, 4, 5})
		expected := []int{0, 9}
		checkSums(t, expected, received)
	})
}
