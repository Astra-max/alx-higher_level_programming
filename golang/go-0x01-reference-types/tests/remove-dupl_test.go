package tests

import (
	"reference/data"
	"reference/libs"
	"reflect"
	"testing"
)

func Test_RemoveDuplic(t *testing.T) {

	for _, tt := range data.DuplicateCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := libs.RemoveDuplic(tt.Input)

			if !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("Expected %v but got %v\n", tt.Expected, got)
			}
		})
	}
}