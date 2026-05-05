import { currentDate } from "./create-date.js";

/**
 * 
 * @returns returns string rep of date
 */

function toString() {
    return currentDate().toString()
}

function getYear() {
    return currentDate().getFullYear()
}

function getMonth() {
    return currentDate().getMonth()
}

function getDay() {
    return currentDate().getDay()
}

function toLocal() {
    return currentDate().toLocaleDateString()
}

function print() {
    console.log(`${toLocal()}`)
}

function toJson() {
    return currentDate().toJSON()
}
print()