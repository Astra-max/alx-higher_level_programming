function varide(...commands: string[]){
    for (let cmd of commands) {
        switch (cmd) {
            case "ls":
                console.log("listing files")
                break
            case "mkdir":
                console.log("creating directory")
                break
            default:
                console.log(cmd, "command not recognized")
        }
    }
}

varide("ls", "mkdir", "pwd", "cd")