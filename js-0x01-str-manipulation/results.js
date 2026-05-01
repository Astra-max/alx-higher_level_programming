function Header(title, padding, pattern) {
    if (typeof title !== "string") {
        throw new Error("Expected string type")
    }
    return title.padStart(padding, pattern)
}

function getAllResults() {
    const examResults = {
        "jame": { math: "A", English: "C", Kiswahili: "C", Science: "A" },
        "Evans": { math: "A", English: "C", Kiswahili: "C", Science: "A" },
        "Joseph": { math: "A", English: "C", Kiswahili: "C", Science: "A" },
        "Eunice": { math: "A", English: "C", Kiswahili: "C", Science: "A" },
    }

    for (let [student, results] of Object.entries(examResults)) {
        console.log(`Name: ${student}\nResults: Math --> ${results.math} English -----> ${results.English} Kiswahili -----> ${results.Kiswahili} Science ----> ${results.Science}\n`)
    }
}

function mySchool() {
    console.log(Header("Astra academy", 25, "==========>"))
    getAllResults()
}
mySchool()