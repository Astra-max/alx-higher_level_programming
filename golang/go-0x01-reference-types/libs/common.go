package libs

func FindCommon(a, b []int) []int {
	results := []int{}

	for i:=0; i<len(a); i++ {
		for j:=0; j<len(b); j++ {
			if a[i] == b[j] {
				results = append(results, a[i])
			}
		}
	}
	return results
}