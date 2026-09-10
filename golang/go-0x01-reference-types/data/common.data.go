package data

type CommonElementsTestModel struct {
	Name     string
	InputA   []int
	InputB   []int
	Expected []int
}

var CommonElementsTestCases = []CommonElementsTestModel{
	{
		Name:     "Find common elements",
		InputA:   []int{1, 2, 3, 4, 5},
		InputB:   []int{3, 5, 7, 9},
		Expected: []int{3, 5},
	},
	{
		Name:     "Multiple common elements",
		InputA:   []int{10, 20, 30, 40},
		InputB:   []int{20, 40, 60},
		Expected: []int{20, 40},
	},
	{
		Name:     "No common elements",
		InputA:   []int{1, 2, 3},
		InputB:   []int{4, 5, 6},
		Expected: []int{},
	},
	{
		Name:     "All elements are common",
		InputA:   []int{1, 2, 3},
		InputB:   []int{1, 2, 3},
		Expected: []int{1, 2, 3},
	},
	{
		Name:     "Handle negative numbers",
		InputA:   []int{-1, -2, -3, 4},
		InputB:   []int{-3, -1, 5, 6},
		Expected: []int{-1, -3},
	},
	{
		Name:     "Handle zero",
		InputA:   []int{0, 1, 2, 3},
		InputB:   []int{0, 3, 5},
		Expected: []int{0, 3},
	},
	{
		Name:     "First slice is empty",
		InputA:   []int{},
		InputB:   []int{1, 2, 3},
		Expected: []int{},
	},
	{
		Name:     "Second slice is empty",
		InputA:   []int{1, 2, 3},
		InputB:   []int{},
		Expected: []int{},
	},
}