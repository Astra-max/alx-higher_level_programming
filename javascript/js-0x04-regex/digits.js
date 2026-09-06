function digits(val) {
	if (val === undefined) return [val]
	let pattern = /(\d+)/g
	return val.match(pattern) || []
}

console.log(digits("dongo67"))
