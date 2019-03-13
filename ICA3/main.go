package main

import (
	"fmt"

	"./fileutils"
)

//Ved å kjøre main via fileutils.go får vi en "byteslice" av de 2 txt filene som vi kan sammenligne.

func main() {

	fmt.Println()

	fmt.Println("ICA3 - Oppgave 1a)")

	fmt.Println("Dette er Text 1")
	fmt.Println(fileutils.FileToByteslice("text1.txt"))

	fmt.Println()

	fmt.Println("Dette er Text 2")
	fmt.Println(fileutils.FileToByteslice("text2.txt"))

	fmt.Println()
}
