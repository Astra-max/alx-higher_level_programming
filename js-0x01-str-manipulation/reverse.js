import { split } from "./split.js"

function reverse(str) {
    let chars = split(str, "")

    let left = 0
    let right = chars.length - 1

    while (left < right) {

        let temp = chars[left]
        chars[left] = chars[right]
        chars[right] = temp

        left++
        right--
    }

    return chars.join("")
}

console.log(reverse("hello"))
