package tests

import (
	"hello/data"
	"hello/libs"
	"testing"
)

func Test_Temperature(t *testing.T) {
	for _, tt := range data.TempTestData {
		t.Run(tt.Name, func(t *testing.T) {
			farent := libs.ConvertToFarenheit(tt.Celcius)

			if farent != tt.Expected {
				t.Fatalf("Exected %v but got %v\n", tt.Expected, farent)
			}
		})
	}
}
