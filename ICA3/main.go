package main

import (
	"fmt"
	"os"

	"./fileutils"
	"./lineshift"
)

//Ved å kjøre main via fileutils.go får vi en "byteslice" av de 2 txt filene som vi kan sammenligne.
//I tillegg sjekker sjekkes hvilken byte for linjeskift som er brukt i de 2 tekst- filene.
func main() {

	fmt.Println()

	fmt.Println("ICA3 - Oppgave 1a)")

	fmt.Println("Dette er Text 1")
	fmt.Println(fileutils.FileToByteslice("text1.txt"))

	fmt.Println()

	fmt.Println("Dette er Text 2")
	fmt.Println(fileutils.FileToByteslice("text2.txt"))

	fmt.Println()

	fmt.Println("ICA3 - Oppgave 1b")
	filename := os.Args[1]
	lineshift.LineShift(filename)
}
