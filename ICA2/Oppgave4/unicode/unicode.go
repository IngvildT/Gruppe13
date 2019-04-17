package unicode

import (
	"fmt"
)

//OPPGAVE 4A
func Translate(expression string, language string) string {
	unicode := ""
	if expression == "nord og sør" {
		if language == "is" {
			unicode = "\x22" + expression + "\x22" + " på islandsk er: " + "\x22\x6E\x6f\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\x22"
		} else if language == "jp" {
			unicode = "\x22" + expression + "\x22" + " på japansk er: " + "\x22\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\x22"
		}
	}
	return unicode
}

//OPPGAVE 4B
func UnicodeCodePointDemo() {
	// Er det likt på MS Windows og macOS?
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	// Demonstrerer at deler av et tegn representert med flere bytes
	// ikke kan tolkes innenfor koden (unicode)
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
