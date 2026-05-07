/**
 * array-like object can't access method to Array(Array.from)
 * soln - convert them to arrays
 * soln2-- function.Array.prototype.forEach.call - doesnt convert the obj
 */


const user = {
    '0': 'max',
    '1': 'john',
    length: 2
}

// length property  -- has to be provided
// converting to array

const userList = Array.from(user)

userList.forEach((val) => console.log(user));

//  calling array methods without converting to an array
// Array.prototype.forEach.call(obj, function)

Array.prototype.forEach.call(user, (user) => console.log(user))