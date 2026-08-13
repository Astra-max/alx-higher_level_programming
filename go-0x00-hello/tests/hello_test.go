package tests

import (
	"hello/libs"
	"testing"
)

func Test_Hello(t *testing.T) {
	got := libs.GoStrs()
	expect := "hello astra!"

	if expect != got {
		t.Fatalf("Expected %v but got %v\n", expect, got)
	}
}
