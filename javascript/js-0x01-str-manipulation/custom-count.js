function count(str1, str2) {
	let count = 0

	let newStr = str1.split(" ")
	for (let match of newStr) {
		if (match === str2) {
			count++
		}
	}
	return count
}

console.log(count("astra max astra", "astra"))
