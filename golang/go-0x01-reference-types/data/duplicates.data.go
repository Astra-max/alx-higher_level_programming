package data


var DuplicateCases = []struct {
	Name     string
	Input    []int
	Expected []int
}{
	{
		Name:     "Remove duplicates",
		Input:    []int{4, 2, 4, 1, 2, 7, 1, 9},
		Expected: []int{4, 2, 1, 7, 9},
	},
	{
		Name:     "All duplicates",
		Input:    []int{5, 5, 5, 5},
		Expected: []int{5},
	},
	{
		Name:     "No duplicates",
		Input:    []int{1, 2, 3, 4, 5},
		Expected: []int{1, 2, 3, 4, 5},
	},
	{
		Name:     "Negative numbers",
		Input:    []int{-1, 3, -1, 5, 3, -7},
		Expected: []int{-1, 3, 5, -7},
	},
	{
		Name:     "Empty slice",
		Input:    []int{},
		Expected: []int{},
	},
}