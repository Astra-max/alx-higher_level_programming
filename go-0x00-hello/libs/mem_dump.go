package libs


func MemDump(data []byte) string {

	hex := "0123456789abcdef"

	results := ""

	for i, byt := range data {
		lower := byt / 16
		upperBound := byt % 16

		if i < len(data) -1 {
			results += string(hex[lower]) + string(hex[upperBound]) + " "
		} else {
			results += string(hex[lower]) + string(hex[upperBound])
		}
	}
	return results
}