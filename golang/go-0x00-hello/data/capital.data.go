package data

type CapitalizeNextModel struct {
	Name     string
	Value    string
	Expected string
}

var CapitalizeNextCases = []CapitalizeNextModel{

	// Empty
	{
		Name:     "empty string",
		Value:    "",
		Expected: "",
	},

	// Single characters
	{
		Name:     "single lowercase letter",
		Value:    "a",
		Expected: "a",
	},

	{
		Name:     "single uppercase letter",
		Value:    "A",
		Expected: "A",
	},

	{
		Name:     "single number",
		Value:    "1",
		Expected: "1",
	},

	{
		Name:     "single special character",
		Value:    "!",
		Expected: "!",
	},

	// Lowercase
	{
		Name:     "lowercase word",
		Value:    "hello",
		Expected: "HELLo",
	},

	{
		Name:     "two lowercase letters",
		Value:    "ab",
		Expected: "Ab",
	},

	{
		Name:     "alphabet",
		Value:    "abcdefghijklmnopqrstuvwxyz",
		Expected: "ABCDEFGHIJKLMNOPQRSTUVWXYz",
	},

	// Uppercase
	{
		Name:     "uppercase word",
		Value:    "HELLO",
		Expected: "HELLO",
	},

	{
		Name:     "uppercase A",
		Value:    "A",
		Expected: "A",
	},

	{
		Name:     "uppercase Z",
		Value:    "Z",
		Expected: "Z",
	},

	// Mixed case
	{
		Name:     "mixed case",
		Value:    "hElLo",
		Expected: "HELLo",
	},

	{
		Name:     "alternating case",
		Value:    "aAaAa",
		Expected: "AAAaA",
	},

	// Numbers
	{
		Name:     "single digit",
		Value:    "5",
		Expected: "5",
	},

	{
		Name:     "all digits",
		Value:    "0123456789",
		Expected: "0123456789",
	},

	{
		Name:     "number followed by letter",
		Value:    "1a",
		Expected: "1a",
	},

	{
		Name:     "letter followed by number",
		Value:    "a1",
		Expected: "A1",
	},

	{
		Name:     "letters and numbers",
		Value:    "abc123",
		Expected: "ABC123",
	},

	// Special characters
	{
		Name:     "special characters",
		Value:    "!@#$%^&*()",
		Expected: "!@#$%^&*()",
	},

	{
		Name:     "symbols between letters",
		Value:    "a-b",
		Expected: "A-B",
	},

	{
		Name:     "underscore",
		Value:    "a_b",
		Expected: "A_B",
	},

	{
		Name:     "slash",
		Value:    "a/b",
		Expected: "A/B",
	},

	{
		Name:     "dot",
		Value:    "a.b",
		Expected: "A.B",
	},

	// Spaces
	{
		Name:     "single space",
		Value:    " ",
		Expected: " ",
	},

	{
		Name:     "word with space",
		Value:    "hello world",
		Expected: "HELLO WORLd",
	},

	{
		Name:     "leading space",
		Value:    " hello",
		Expected: " HELLo",
	},

	{
		Name:     "trailing space",
		Value:    "hello ",
		Expected: "HELLO ",
	},

	{
		Name:     "multiple spaces",
		Value:    "hello   world",
		Expected: "HELLO   WORLd",
	},

	// Punctuation
	{
		Name:     "comma",
		Value:    "hello,world",
		Expected: "HELLO,WORLd",
	},

	{
		Name:     "period",
		Value:    "hello.world",
		Expected: "HELLO.WORLd",
	},

	{
		Name:     "exclamation",
		Value:    "hello!world",
		Expected: "HELLO!WORLd",
	},

	{
		Name:     "question mark",
		Value:    "hello?world",
		Expected: "HELLO?WORLd",
	},

	// Boundaries
	{
		Name:     "lowercase a",
		Value:    "a",
		Expected: "a",
	},

	{
		Name:     "lowercase z",
		Value:    "z",
		Expected: "z",
	},

	{
		Name:     "uppercase letter followed by lowercase",
		Value:    "Ab",
		Expected: "Ab",
	},

	{
		Name:     "lowercase followed by uppercase",
		Value:    "aB",
		Expected: "AB",
	},

	{
		Name:     "digit zero followed by letter",
		Value:    "0a",
		Expected: "0a",
	},

	{
		Name:     "digit nine followed by letter",
		Value:    "9a",
		Expected: "9a",
	},

	{
		Name:     "uppercase and number",
		Value:    "A1",
		Expected: "A1",
	},

	{
		Name:     "number and uppercase",
		Value:    "1A",
		Expected: "1A",
	},

	// Mixed
	{
		Name:     "mixed letters numbers symbols",
		Value:    "a1-b2_c3",
		Expected: "A1-B2_C3",
	},

	{
		Name:     "email like string",
		Value:    "hello@example.com",
		Expected: "HELLO@EXAMPLE.COm",
	},

	{
		Name:     "url like string",
		Value:    "https://example.com",
		Expected: "HTTPS://EXAMPLE.COm",
	},

	// Unicode
	{
		Name:     "unicode lowercase",
		Value:    "é",
		Expected: "é",
	},

	{
		Name:     "unicode uppercase",
		Value:    "É",
		Expected: "É",
	},

	{
		Name:     "emoji",
		Value:    "😀",
		Expected: "😀",
	},

	{
		Name:     "unicode mixed with ascii",
		Value:    "aéz",
		Expected: "AéZ",
	},
}
