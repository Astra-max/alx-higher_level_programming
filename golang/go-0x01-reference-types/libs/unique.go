package libs

func UniqueNumbers(arr [6]int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	unique := make(map[int]bool)
	results := []int{}


	for _, num := range arr {
		if _,ok := unique[num]; !ok {
			unique[num] = true
			results = append(results, num)
		}
	}

	return results
}