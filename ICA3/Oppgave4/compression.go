package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave4/pipe"
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
	fmt.Println("Første ord original tekst: ", fil[:20])
	fmt.Println("Antall bytes i fil opprinnelig: " + filTotalt + "\n")

	// Henter Hex funksjonen fra pipe, teller antall bytes totalt og printer de 10 første.
	hexString := pipe.Hex(fil)
	hexTotalt := fmt.Sprintf("%d", len(hexString))
	//I hexadesimal representasjon er 1 karakter 4 bits, derav 12 bytes er 24 karakterer hex
	hexSlice := hexString[:20]

	fmt.Println("Antall bytes totalt: " + hexTotalt + "\n" +
		"20 første bytes Hexadesimal: " + hexSlice + "\n")

	// Henter base64 funksjonen fra pipe, teller antall bytes totalt og printer de 10 første.
	base64String := pipe.Base64(fil)
	base64Totalt := fmt.Sprintf("%d", len(base64String))
	//I base64 format er 1 karakter 6 bits, dvs for å få 12 bytes trengs det 16 karakterer base64
	base64Slice := base64String[:20]
	base64Hex := fmt.Sprintf("% X", base64String[:20])

	fmt.Println("Antall bytes totalt: " + base64Totalt + "\n" +
		"20 første bytes av Hex i base64 format: " + base64Slice + "\n" + "Base64 konvertert til Hex: " + base64Hex + "\n")

	// Henter Gzip funksjonen fra pipe, teller antall bytes totalt og printer de 10 første.
	gzipString := pipe.Gzip(fil)
	gzipTotalt := fmt.Sprintf("%d", len(gzipString))
	gzipSlice := fmt.Sprintf("%s", gzipString[:20])
	//Konvertering til Hexadesimal etter komprimering
	gzipHex := fmt.Sprintf("% X", gzipString[:20])

	fmt.Println("Antall bytes totalt: " + gzipTotalt + "\n" +
		"20 første bytes av base64 i Gzip format: " + gzipSlice + "\n" + "Gzip konvertert til Hex: " + gzipHex + "\n")

}
