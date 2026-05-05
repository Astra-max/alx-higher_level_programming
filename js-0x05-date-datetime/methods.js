import { currentDate } from "./create-date";

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