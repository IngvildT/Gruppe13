package main

import (
	"fmt"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA2/Oppgave2/iso"
)

func main() {
	iso.IterateOverExtendedASCIIStringLiteral(iso.AsciiExtended)
	fmt.Printf(iso.GreetingExtendedASCII())
}
