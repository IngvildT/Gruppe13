package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA3/Oppgave4/pipe"
)

//OPPGAVE 4C compression
//https://gobyexample.com/command-line-arguments

func main() {
	filNavn := os.Args[1]
	file, err := ioutil.ReadFile(filNavn)
	//Åpner filen (parameteret filNavn) fra kommandolinja og leser den.

	if err != nil {
		fmt.Println(err)
	} //Hvis åpningen av filen failer printes en errormelding.

	fil := string(file)
	filTotalt := fmt.Sprintf("%d", len(fil))
	fmt.Println("Antall bytes i fil opprinnelig: " + filTotalt + "\n")

	// Henter Hex funksjonen fra pipe, teller antall bytes totalt og printer de 20 første.
	hexString := pipe.Hex(fil)
	hexTotalt := fmt.Sprintf("%d", len(hexString))
	hexSlice := hexString[:20]

	fmt.Println("Antall bytes totalt: " + hexTotalt + "\n" +
		"Innhold 20 første bytes: " + hexSlice + "\n")

	// Henter base64 funksjonen fra pipe, teller antall bytes totalt og printer de 20 første.
	base64String := pipe.Base64(fil)
	base64Totalt := fmt.Sprintf("%d", len(base64String))
	base64Slice := base64String[:20]

	fmt.Println("Antall bytes totalt: " + base64Totalt + "\n" +
		"Innhold 20 første bytes: " + base64Slice + "\n")

	// Henter Gzip funksjonen fra pipe, teller antall bytes totalt og printer de 20 første.
	gzipString := pipe.Gzip(fil)
	gzipTotalt := fmt.Sprintf("%d", len(gzipString))
	gzipSlice := gzipString[:20]

	fmt.Println("Antall bytes totalt: " + gzipTotalt + "\n" +
		"Innhold 20 første bytes: " + gzipSlice + "\n")

}
