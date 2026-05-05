const userData = {name: "astra", age: 89}

const clone = Object.assign({}, userData) // object.assign
const clone1 = {...userData } // es6 spread operator
const copy = userData //passing ref\\

userData.age += 23

console.log(copy, clone, clone1)