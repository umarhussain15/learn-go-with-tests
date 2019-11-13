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

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rect := Rectangle{10, 20}
		checkArea(t, rect, 200.0)
	})
	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, math.Pi*100)
	})

}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, math.Pi * 100},
	}

	for _, v := range areaTests {
		got := v.shape.Area()
		if got != v.want {
			t.Errorf("got %g want %g", got, v.want)
		}
	}
}
