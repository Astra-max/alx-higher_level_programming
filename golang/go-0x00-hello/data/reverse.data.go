package data

type ReverseModel struct {
	TestName string
	Value    string
	Expected string
}

var ReverseStrData []ReverseModel = []ReverseModel{
	{
		TestName: "simple word",
		Value:    "hello",
		Expected: "olleh",
	},
	{
		TestName: "another word",
		Value:    "world",
		Expected: "dlrow",
	},
	{
		TestName: "palindrome",
		Value:    "level",
		Expected: "level",
	},
	{
		TestName: "single character",
		Value:    "a",
		Expected: "a",
	},
	{
		TestName: "empty string",
		Value:    "",
		Expected: "",
	},
	{
		TestName: "two characters",
		Value:    "go",
		Expected: "og",
	},
	{
		TestName: "sentence",
		Value:    "hello world",
		Expected: "dlrow olleh",
	},
	{
		TestName: "numbers",
		Value:    "12345",
		Expected: "54321",
	},
	{
		TestName: "mixed characters",
		Value:    "Go123",
		Expected: "321oG",
	},
	{
		TestName: "spaces",
		Value:    "go test",
		Expected: "tset og",
	},
}
