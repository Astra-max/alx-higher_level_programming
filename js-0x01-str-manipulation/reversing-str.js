let name = "james"

let reversedName = name.split('').reverse().join('')
console.log(reversedName)

function reverseName(name) {
    let rev = ""
    for (let i=name.length -1; i>=0; i--) {
        rev += name[i]
    }
    console.log(rev)
}
reverseName(name)