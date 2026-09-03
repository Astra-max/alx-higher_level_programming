package libs


func ReverseStr(str string) string {
	if len(str) == 0 {
		return ""
	}
	runeStr := []rune(str)

	for i,j := 0, len(runeStr)-1; i<j; i,j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}