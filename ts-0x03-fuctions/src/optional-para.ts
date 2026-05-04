function playing(music: string, artist?: string) {
    if (music && artist) {
        console.log(`playing ${music} by ${artist}`)
        return
    }
    console.log(`playing music ${music}`)
}
playing("tattoo")
playing("tatoo", "loreen")