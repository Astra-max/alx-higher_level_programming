package tests

import (
	"hello/data"
	"hello/libs"
	"testing"
)

func Test_ReverseStr(t *testing.T) {
	for _,tt := range data.ReverseStrData {
		t.Run(tt.TestName, func(t *testing.T) {
			got := libs.ReverseStr(tt.Value)

			if got != tt.Expected {
				t.Fatalf("Expected %v but got %v\n", tt.Expected, got)
			}
		})
	}
}