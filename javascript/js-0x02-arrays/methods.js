export const ages = [2,3,4,5,6,4,1]

const totalAge = ages.reduce((a,b)=> a+b) // summing all values in arrray
// mapping through an array returns new modified array

console.log(`Total users ages: ${totalAge}`)
