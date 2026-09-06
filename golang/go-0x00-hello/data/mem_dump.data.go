package data

type MemDumpModel struct {
	Name     string
	Value    []byte
	Expected string
}

var MemDumpCases []MemDumpModel = []MemDumpModel{
	{
		Name:     "empty input",
		Value:    []byte{},
		Expected: "",
	},
	{
		Name:     "single zero byte",
		Value:    []byte{0},
		Expected: "00",
	},
	{
		Name:     "single byte",
		Value:    []byte{65},
		Expected: "41",
	},
	{
		Name:     "single byte zero padding",
		Value:    []byte{10},
		Expected: "0a",
	},
	{
		Name:     "single byte maximum",
		Value:    []byte{255},
		Expected: "ff",
	},
	{
		Name:     "multiple bytes",
		Value:    []byte{65, 66, 67},
		Expected: "41 42 43",
	},
	{
		Name:     "zero values",
		Value:    []byte{0, 0, 0},
		Expected: "00 00 00",
	},
	{
		Name:     "maximum values",
		Value:    []byte{255, 255, 255},
		Expected: "ff ff ff",
	},
	{
		Name:     "mixed boundaries",
		Value:    []byte{0, 1, 15, 16, 17, 31, 32, 255},
		Expected: "00 01 0f 10 11 1f 20 ff",
	},
	{
		Name:     "hex digits lower range",
		Value:    []byte{10, 11, 12, 13, 14, 15},
		Expected: "0a 0b 0c 0d 0e 0f",
	},
	{
		Name:     "hex digits upper range",
		Value:    []byte{160, 171, 188, 205, 222, 239},
		Expected: "a0 ab bc cd de ef",
	},
	{
		Name:     "hello",
		Value:    []byte("Hello"),
		Expected: "48 65 6c 6c 6f",
	},
	{
		Name:     "numbers as bytes",
		Value:    []byte("12345"),
		Expected: "31 32 33 34 35",
	},
	{
		Name:     "spaces",
		Value:    []byte("hello world"),
		Expected: "68 65 6c 6c 6f 20 77 6f 72 6c 64",
	},
	{
		Name:     "newline and tab",
		Value:    []byte{'\n', '\t'},
		Expected: "0a 09",
	},
	{
		Name:     "all important boundaries",
		Value:    []byte{15, 16, 31, 32, 127, 128, 254, 255},
		Expected: "0f 10 1f 20 7f 80 fe ff",
	},
}
