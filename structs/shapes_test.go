package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10, 20}
	got := Perimeter(rect)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("Rectangles", func(t *testing.T) {
		rect := Rectangle{10, 20}
		got := rect.Area()
		want := 200.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := math.Pi * 100

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
