package tests

import (
	"testing"

	"hello/data"
	"hello/libs"
)

func TestCapitalizeNext(t *testing.T) {
	for _, tt := range data.CapitalizeNextCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := libs.CapitalizeNext(tt.Value)

			if got != tt.Expected {
				t.Errorf(
					"CapitalizeNext(%q) = %q, expected %q",
					tt.Value,
					got,
					tt.Expected,
				)
			}
		})
	}
}
