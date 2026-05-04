/**
 * optional/default values should be passed at end of function param list
 * @param music 
 * @param artist - optional param
 * @returns -- void
 */



function playing(music: string, artist?: string) {
    if (music && artist) {
        console.log(`playing ${music} by ${artist}`)
        return
    }
    console.log(`playing music ${music}`)
}
playing("tattoo")
playing("tatoo", "loreen")