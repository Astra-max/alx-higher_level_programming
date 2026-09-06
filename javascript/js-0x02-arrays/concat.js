import { ages } from "./methods.js"
import { friends } from "./looping.js"

const newFriendsAgs = [...friends, ...ages] // via spread operator recommended

// via array.concat method

const concatFriends = ages.concat(friends)

console.log(newFriendsAgs)