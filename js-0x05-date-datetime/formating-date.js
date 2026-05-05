import { currentDate } from "./create-date.js";

/**
 * date -- toLocaleDateString(local, {options})
 */

function format() {
    const options = {day: 'numeric', month: 'short', year: 'numeric'}
    const formated = currentDate().toLocaleDateString("en-GB", options)
    return formated
}

function formatTime() {
    const light = currentDate().getHours()
    const meridiam = light > 12 ? "am" : "pm"
    const options = { hour: '2-digit', minute: '2-digit'}
    return currentDate().toLocaleDateString("en-GB", options) + " " + meridiam
}

console.log(formatTime())