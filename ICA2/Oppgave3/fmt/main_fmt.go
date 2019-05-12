package main

import (
	"fmt"
)

//OPPGAVE 3A - fmt pakken og verb
func main() {
	
	//Bytesekvens 1 fra oppgaveteksten ([½ ² = ¼   â ​​]")
	bytes1 := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	//Bytesekvens 2 fra oppgaveteksten (½?=? ⌘")
	bytes2 := "\xc2\xbd\x3f\x3d\x3f\xe2\x8c\x98"

	println("\n")
	println("Resultat med %s verb:")
	fmt.Printf("%s\n", bytes1)
	println("\n")

	println("Resultat med %q verb:")
	fmt.Printf("%q\n", bytes1)
	println("\n")

	println("Resultat med %+q verb:")
	fmt.Printf("%+q\n", bytes1)
	println("\n")

	println("Resultat med %c verb:")
	fmt.Printf("%c\n", bytes1)
	println("\n")

	println("Bytesekvens 1⁄2?=? ⌘ med %s verb:")
	println("\n")
	fmt.Printf("%s\n", bytes2)
	println("\n")
}
