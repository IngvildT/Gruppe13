package pipe

// https://www.socketloop.com/references/golang-io-pipe-function-example
// https://blog.golang.org/strings
// https://golang.org/pkg/encoding/base64/#example_Encoding_EncodeToString
// https://medium.com/go-walkthrough/go-walkthrough-fmt-55a14bbbfc53

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"log"
)

////OPPGAVE 4C

//Funksjon 1: Returnerer en hexadesimal representasjon basert på ASCII/UTF-8 koding.
func Hex(c string) string {
	fmt.Println("Funksjon 1 - Hexadesimal:")
	/** Siden hexadecimal systemet består av 16 symboler(base16) brukes Verb %X (base 16) for å konvertere 
	s stringen til ASCII/UTF8 koding og printer konverteringen. Siden vi ønsker å holde på resultatet 
	av denne funksjonen brukes Sprintf for å formattere stringen uten å printe til os.Stdout. */

	return fmt.Sprintf("% X", c)
}

//Funksjon 2: Returnerer en hexadesimal funksjon basert på base64.
func Base64(c string) string {
	fmt.Println("Funksjon 2 - base64:")
	//Konverterer s string fra base 16 (%X) til base 64 via EncodeToString og returnerer konverteringen som en string.
	return fmt.Sprintln(base64.StdEncoding.EncodeToString([]byte(c)))
}

//Funksjon 3: Returnerer en hexadesimal funksjon basert på base64 komprimert ved bruk av gzip.
func Gzip(c string) string {
	fmt.Println("Funksjon 3 - base64 komprimert med gzip:")
	//Åpner en buffer og Writer som komprimerer s string base 64 til gzip og printer tilslutt konverteringen.
	var buffer bytes.Buffer
	gZipWriter := gzip.NewWriter(&buffer)
	//Skriver s tring inni bufferen, dersom problemer gis en feilmelding.
	_, err := gZipWriter.Write([]byte(c))
	if err != nil {
		log.Fatal(err)
	}
	//Lukker skriveren
	err = gZipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
	//Returnerer innholdet i bufferen som en string. 
	return fmt.Sprintln(buffer, "% X")
}


