package tests

import (
	"reference/data"
	"reference/libs"
	"reflect"
	"testing"
)

func Test_Unique(t *testing.T) {
	for _, tt := range data.UniqueNumbersTestCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := libs.UniqueNumbers(tt.Input)

			if !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("Expected %v but got %v\n", tt.Expected, got)
			}
		})
	}
}

func Test_Duplicates(t *testing.T) {}