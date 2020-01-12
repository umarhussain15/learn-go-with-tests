package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Umar")

	got := buffer.String()
	want := "Hello, Umar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
