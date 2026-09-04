package libs

func CapitalizeNext(str string) string {
	results := ""
	capNext := false

	for _, char := range str {
		if capNext {
			results += string(Capitalize(char))

			capNext = false
			continue
		}

		if IsNum(char) {
			capNext = true
		}
	}
	return results
}

func IsNum(data rune) bool {
	return data >= '0' && data <= '9'
}

func Capitalize(data rune) rune {
	if data >= 'a' && data <= 'z' {
		return data - 32
	}
	return data
}
