package lineshift

import (
	"fmt"
	"strings"
	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave1/fileutils"
)
//Boolean funksjonen LineShift sjekker om filen inneholder CR "\r" og LF "\n".

func LineShift(filnavn string) {
	byteSlice := fileutils.FileToByteslice(filnavn)

	s := string(byteSlice[:])

	fmt.Print("Er det brukt bytes for Carriage Return (CR) i denne filen?: ")
	fmt.Println(strings.Contains(s, "\r"))

	fmt.Print("Er det brukt bytes for Line Feed (LF) i denne filen?: ")
	fmt.Println(strings.Contains(s, "\n"))

}
