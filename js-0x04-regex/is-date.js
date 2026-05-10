function isDate(val) {
	const pattern = /\d{2}-\d{2}-\d{4}/g
	return val.match(pattern) || []
}

console.log(isDate("today 10-02-2026 independece 01-06-1964"))
