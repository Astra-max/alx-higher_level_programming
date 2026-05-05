const user = {name: 'max'};

Object.defineProperty(user, 'age', {get: ()=>89, set: (value)=>value})
console.log(user.age)