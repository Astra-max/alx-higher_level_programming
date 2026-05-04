type User = {
    name: string
    age: number
    dob: string
}

function saveUser({name, age}: User) {
    console.log(`saving ${name} of age ${age} years old`)
}

saveUser({name: "max", age: 23, dob: ""})