function compare(a, b) {
    if (a === b) {
        return 0
    } else if (a > b) {
        return 1
    } else {
        return -1
    }
}

// sorting a string slice

let friends = ['james', 'okello', 'ahn', 'oketch']

let sorted = friends.sort((a,b)=>compare(a,b))
console.log(sorted)


//javascripts stores strings as one code unit
//surrogate char like emoji requires two code units

let test = "\uD800".isWellFormed()
console.log(test)

// methods
// endsWith
//includes
//indexOf
//trimStart
//trimEnd
//replace
//repeat
//padStart
//padEnd