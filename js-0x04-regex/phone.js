function phoneNumber(number) {
	const pattern = /(\+254)?7[0-9]\d{9}/g
	return number.match(patter) || []
}

console.log(phoneNumber("+254789542378"))
