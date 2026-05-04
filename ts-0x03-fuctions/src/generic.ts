/**
 * @function -- generic function that returns generic value
 * @param value -- generic parameter
 * @returns -- return generic parameter
 */


function  getValue<T>(value: T): T {
    return value
}

// i can take anything
console.log(getValue<string>("maxwel"))
console.log(getValue<number>(56))
console.log(getValue(true))
