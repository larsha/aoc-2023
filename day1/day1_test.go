package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	n := Part1()

	want := 54667

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}

func TestPart2(t *testing.T) {
	n := Part2()

	want := 54203

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}
