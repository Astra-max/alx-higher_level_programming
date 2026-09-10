package libs

func FindCommon(a, b []int) []int {
	results := []int{}

	for i:=0; i<=len(a)-1; i++ {
		for j:=i+1; j<=len(b)-1; j++ {
			if a[i] == a[j] {
				results = append(results, a[i])
			}
		}
	}
	return results
}