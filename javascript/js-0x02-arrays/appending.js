import { ages } from "./methods.js"

// adding infrom of the array

const newAges = ages.unshift(23, 45) // return is new length of array
ages.push(30) // appends value at the end of list

ages.pop() // remove from end return the array
ages.shift() // removes first element from the array and return the array

// inserting at a particular inder

ages.splice(1, 0, 79, 90)