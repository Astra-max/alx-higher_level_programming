let isActive = true
let completedTest = true

const color = isActive? "green": "gray"

console.log("User active color: ", color)

/**
 * tenary operator - can also e nested like nested ifs
 * 
 */

const colorWithTest = isActive ? (completedTest? "blue" : "green") : "gray"
console.log("completed tests color: ",colorWithTest)