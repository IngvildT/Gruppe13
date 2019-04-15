package main

import (
	"fmt"

	speech "github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA4/Oppgave3/wit/go-speak"
)

//example

func main() {
	speech.SetWitKey("NQKGDPHCW465PKZIAWQ3D4RM5C5KCGYE") //Wit API Key MUST be set before calling any other Wit.AI functions
	a := speech.SendWitVoice("./audio.wav")
	fmt.Println(a)
}
