function optional(str) {
	const pattern = /colou?r/g
	return str.match(pattern) || []
}

console.log(optional("colour or color both work for me"))
