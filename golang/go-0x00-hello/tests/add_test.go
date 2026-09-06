package tests


import (
	"testing"

	"hello/libs"
)

func Test_Add(t *testing.T) {
	expected := 68
	got := libs.Add(24, 44)

	if expected != got {
		t.Fatalf("Expected %v got %v\n", expected, got)
	}
}
