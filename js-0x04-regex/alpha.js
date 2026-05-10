const upperCase = str => str.match(/[A-Z]/) || []
const lowerCase = str => str.match(/[a-z]/g) || []
console.log(upperCase("Astra").length > 0)
