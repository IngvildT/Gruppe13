package main

import (
	"fmt"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA3/Oppgave4/pipe"
)

func main() {

	hex := pipe.Hex("Ulan Bator er hovedstaden i Mongolia")
	fmt.Println(hex, "\n")

	b64 := pipe.Base64(hex)
	fmt.Println(b64, "\n")

	gzip := pipe.Gzip(b64)
	fmt.Println(gzip)
}
