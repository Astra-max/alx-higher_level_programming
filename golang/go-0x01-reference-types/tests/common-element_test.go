package tests

import (
	"reference/data"
	"reference/libs"
	"reflect"
	"testing"
)

func Test_CommonElement(t *testing.T) {

	for _, tt := range data.CommonElementsTestCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := libs.FindCommon(tt.InputA, tt.InputB)

			if !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("Expected %v but got %v\n", tt.Expected, got)
			}
		})
	}
}