package data

type TestModel struct {
	Name     string
	Num1     int
	Num2     int
	Expected int
}

var AddTestCases []TestModel = []TestModel{
	{Name: "Add 5 and 5", Num1: 5, Num2: 5, Expected: 10},
	//{ Name: "Addition resulting to overflow", Num1:  int(^uint(0) >> 1), Num2: 1, Expected:  9_223_372_036_854_775_808},
	{Name: "Add 5 and 5", Num1: 5, Num2: 5, Expected: 10},
	{Name: "Add 5 and 5", Num1: 5, Num2: 5, Expected: 10},
}
