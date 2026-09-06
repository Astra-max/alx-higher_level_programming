import { ages } from "./methods.js"

const found = ages.find((value) => value === 4)
// returns the value if exists or undefined if not exist

const byIndex = ages.findIndex((v, i) => i === 90) // return index if found or -1
const byLast = ages.findLast()
const byLastIndex = ages.findLastIndex()
console.log(found, byIndex)