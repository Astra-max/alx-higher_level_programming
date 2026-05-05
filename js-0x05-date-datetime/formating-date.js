import { currentDate } from "./create-date.js";

/**
 * date -- toLocaleDateString(local, {options})
 */

function format() {
    const options = {day: 'numeric', month: 'short', year: 'numeric'}
    const formated = currentDate().toLocaleDateString("en-GB", options)
    return formated
}

console.log(format())