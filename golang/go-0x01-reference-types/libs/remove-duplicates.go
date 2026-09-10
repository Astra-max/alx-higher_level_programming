package libs


func RemoveDuplic(arr []int) []int {
	results := []int{}

	seen := make(map[int]bool)

	for _, num := range arr {
		if _,visited := seen[num]; !visited {
			results = append(results, num)
			seen[num] = true
		}
	}
	return results
}