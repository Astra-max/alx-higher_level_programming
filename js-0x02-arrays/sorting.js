import { ages } from "./methods.js"

/**
 * sorting - in a particular order
 */

const asc = (a,b) => a - b;

const newlySorted = ages.sort(asc)
const descOrder = ages.reverse()
console.log(descOrder)
console.log(newlySorted)