package data

type UniqueNumbersTestModel struct {
	Name     string
	Input    [6]int
	Expected []int
}

var UniqueNumbersTestCases = []UniqueNumbersTestModel{
	{
		Name:     "Remove duplicate numbers",
		Input:    [6]int{1, 2, 2, 3, 3, 4},
		Expected: []int{1, 2, 3, 4},
	},
	{
		Name:     "All numbers are unique",
		Input:    [6]int{1, 2, 3, 4, 5, 6},
		Expected: []int{1, 2, 3, 4, 5, 6},
	},
	{
		Name:     "All numbers are the same",
		Input:    [6]int{5, 5, 5, 5, 5, 5},
		Expected: []int{5},
	},
	{
		Name:     "Handle zero values",
		Input:    [6]int{0, 1, 0, 2, 1, 3},
		Expected: []int{0, 1, 2, 3},
	},
	{
		Name:     "Handle negative numbers",
		Input:    [6]int{-1, -2, -1, -3, -2, -4},
		Expected: []int{-1, -2, -3, -4},
	},
	{
		Name:     "Handle positive and negative numbers",
		Input:    [6]int{-1, 1, -1, 2, 1, 2},
		Expected: []int{-1, 1, 2},
	},
}