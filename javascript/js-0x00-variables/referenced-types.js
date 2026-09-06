/**
 * Referenced types
 * plain objest -- {}
 * Arrays -- []
 * functions - () => {}
 * sets -- 
 * maps
 */

let user = {name: "maxwel", age: 45,
    isMarried: false,
    doB: new Date('01-07-2007').toLocaleDateString()
}
let friends = ['james', 'okello', 'clinton']

// maps -- contains key value pairs

let friendsOccupation = new Map()
friendsOccupation.set("jobs", "it")
friendsOccupation.set("games", "chess")

// sets - contains only unique elements
// attempting to add duplicte gets ignored

let months = new Set()
months.add("january")
months.add("febebruary")
months.add("march")


console.log(`${user.name} is ${user.age}yrs old is married ${user.isMarried} dob ${user.doB.toString()}`)
console.log(`My friends are ${friends}`)
console.log("All my friends like playing", friendsOccupation.get("games"))
console.log("Months of the year", months)