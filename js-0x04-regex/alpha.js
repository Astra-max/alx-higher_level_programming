const upperCase = str => str.match(/[A-Z]/g) || []
const lowerCase = str => str.match(/[a-z]/g) || []
const allCase = str => str.match(/[a-z]/gi) || []
console.log(allCase("Astra"))
