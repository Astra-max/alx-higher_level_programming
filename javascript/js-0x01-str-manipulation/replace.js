function replace(str, old, newStr) {
	let newStr1 = str.split(" ")

	for (let i=0; i<newStr1.length; i++) {
		if (newStr1[i] === old) {
			newStr1[i] = newStr
		}
	}
	return newStr1.join(" ")
}

console.log(replace("astra max", "max", "6"))
