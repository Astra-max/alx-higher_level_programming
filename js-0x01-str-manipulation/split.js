function split(str, delim) {
	let results = []
	let startIdx =0
	let delimLen = delim.length-1

	for (let i =0; i<str.length; i++) {
		if (str.slice(i, i + delimLen) === delim) {
			results.push(str.slice(startIdx,i))
			startIdx = delimLen + i
		}
	}
	results.push(str.slice(startIdx))
	return results
}

console.log(split("s_s_s", "_"))
