function Authorise(user: string, next: (msg: string)=>void): ()=>void {
    return () => {
        if (user === "admin") {
            return next("passing request to destination")
        } else {
            return next("unauthorized access")
        }
    }
}

function next(msg: string):void {
    console.log(`${msg}`)
}

Authorise("admin", next)()