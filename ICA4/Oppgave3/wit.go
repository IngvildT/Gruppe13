package main

import (
	"fmt"

	speech "github.com/nicolaifsf/go-speak"
)

//example

func main() {
	speech.SetWitKey("AUQPSPJ4ASSMSDRQQ2Z4FKQGUC5YLSDF") //Wit API Key MUST be set before calling any other Wit.AI functions
	a := speech.SendWitVoice("./audio.wav")
	fmt.Println(a)
}
