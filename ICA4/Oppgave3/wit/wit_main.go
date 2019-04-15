package main

import (
	"fmt"

	speech "github.com/IngvildT/wit/go-speak"
)

//example

func main() {
	speech.SetWitKey("NQKGDPHCW465PKZIAWQ3D4RM5C5KCGYE") //Wit API Key MUST be set before calling any other Wit.AI functions
	a := speech.SendWitVoice("./test.wav")
	fmt.Println(a)
}
