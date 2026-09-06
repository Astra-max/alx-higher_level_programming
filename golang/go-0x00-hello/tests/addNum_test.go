package tests

import (
	"hello/data"
	"hello/libs"
	"testing"
)

func Test_AddNum(t *testing.T) {
	for _, tt := range data.AddTestCases {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := libs.AddNums(tt.Num1, tt.Num2)

			if err != nil {
				t.Fatalf("Got error %v", err)
			}
			if got != tt.Expected {
				t.Fatalf("Expected %v but got %v\n.", tt.Expected, got)
			}
		})
	}
}
