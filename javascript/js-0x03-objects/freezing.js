/**
 * objects are mutable by default
 * making them immutable via object.freeze
 */


const user = {
    friends: ['james', 'julie'],
    hobbies: {
        name: "guiter",
    },
}

// get all object properties including non enumerable
// object.keys --- returns only enumerable
const userProperties = Object.getOwnPropertyNames(user)

for (let key of userProperties) {
    let value = user[key]

    if (typeof value === "object") {
        Object.freeze(value)
    }
}

// freezing nested objects

user.friends[0] = 'emma'

console.log(user.friends[0])

//