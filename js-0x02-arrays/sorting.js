import { ages } from "./methods.js"

/**
 * sorting - in a particular order
 */


// sort -- converts the array to string
// a - b -- forces the values to numbers

const asc = (a,b) => a - b;

const newlySorted = ages.sort(asc)

console.log(newlySorted)
const descOrder = ages.reverse()
console.log(descOrder)

const friends = ['joseph', 'okello', 'oketch', 'adongo', 'wairachu']

/**
 * perfect way of sorting string array 
 * alternative to go character by character sort
 */

const sortStr = (a,b) => a.length - b.length // all fields must be strings

/**
 * sorting via localecompare function
 */

function sortCompare(a, b) {
    if (a > b) return 1
    if (a < b) return -1
    return 0
}

const sortedFriends = friends.sort(sortCompare)
console.log(sortedFriends)