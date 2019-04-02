package main

import (
	"fmt"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA3/Oppgave4/pipe"
)

func main() {

	hex := pipe.Hex("Ulan Bator er hovedstaden i Mongolia.")
	//Printer funksjonen Hex og tar inn stringen c som parameter
	fmt.Println(hex, "\n")

	b64 := pipe.Base64(hex)
	//Printer funksjonen Base64 og tar resultatet av hex funksjonen som paramter.
	fmt.Println(b64, "\n")

	gzip := pipe.Gzip(b64)
	//Printer funksjonen Gzip og tar resultatet av base64 funksjonen som paramter.
	fmt.Println(gzip)
}
