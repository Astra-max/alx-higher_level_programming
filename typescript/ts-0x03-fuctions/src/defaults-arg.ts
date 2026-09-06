/**
 * 
 * @param music 
 * @param artist - is default value that will be used when its value is argument is not
 * passed 
 */


function playMusic(music: string, artist="kenny rodger") {
    console.log(`playing ${music} by ${artist}`)
}

playMusic("city guy")