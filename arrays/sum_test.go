package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 4, 5}
		got := Sum(numbers)
		want := 10
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
