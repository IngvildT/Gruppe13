package main

import (
	"fmt"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA2/Oppgave2/iso"
)

func main() {
	//Kjører funksjonen IterateOverExtendedASCIIStringLiteral
	iso.IterateOverExtendedASCIIStringLiteral(iso.AsciiExtended)
	//Printer funksjonene GreetingExtendedASCII 1 og 2
	fmt.Printf(iso.GreetingExtendedASCII1())
	fmt.Printf(iso.GreetingExtendedASCII2())
}
