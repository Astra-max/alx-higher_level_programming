function coolReturn(name: string): string {
    if (name !== "astra") { 
        name = "james"
    } else if (name === "astra") {
        name = "astra"
    }
    return name
}