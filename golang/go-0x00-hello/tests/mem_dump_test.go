package tests

import (
	"hello/data"
	"hello/libs"
	"testing"
)

func Test_MemDump(t *testing.T) {

	for _, tt := range data.MemDumpCases {
		t.Run(tt.Name, func(t *testing.T) {
			dump := libs.MemDump(tt.Value)

			if dump != tt.Expected {
				t.Errorf("Expected %v but got %v\n", tt.Expected, dump)
			}
		})
	}
}
