/**
 * Strings any data denoted by "", '', or ``
 * double quote, single quote or back tick
 */

let myName = "max"
let nikName = 'astra'
let gitHubUser = `Astra-max`
//let likes = String('food')

console.log("Maxwel in many ways")
console.log(`My real name ${myName}`)
console.log(`My nickname ${nikName}`)
console.log(`My github name ${gitHubUser}`)

//////////////////////////////////
/**
 * string methods
 */

// concatinating strings

let userName = nikName + myName
let userNameByMethod = nikName.concat(myName)
console.log(userName, userNameByMethod)