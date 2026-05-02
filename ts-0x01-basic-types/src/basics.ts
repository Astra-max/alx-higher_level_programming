let userName: string = "astra";
let age: number = 26;
let isMarried: boolean = false;

console.log(`Hello my name is ${userName} am ${age} old and am ${isMarried?"marriade":"not marriade"}`)

/**
 * any - not type safety
 * unknown - check before using it
 * null - variable declared, value set to nothing
 * undefined - variable declared but no value provided
 */


let commingOfChrist: unknown;
let isSick: boolean;
let spaceContains = null;
let holds: any = userName;
holds = age;
commingOfChrist = "";

if (typeof commingOfChrist === "string") {
	commingOfChrist = "soon".toUpperCase();
} else {
	commingOfChrist = "sooner";
}

console.log(`To all christians christ is comming ${commingOfChrist}`);
console.log("Any makes me sick: ", holds);
