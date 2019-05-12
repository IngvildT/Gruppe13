package main

import (
	"fmt"

	slice "Gruppe13/ICA2/Oppgave5/slice/slicekode"
)

func main() {

	// Printer ut funksjonene hvor vi allokerer og deklarere slice
	fmt.Println(slice.AllocateMake(5))
	fmt.Println(slice.AllocateVar(7))

	// Oppretter byteslice1
	// bruker make til å lage en slice og setter inn byte, lidx int, uidx int som parameter.
	// Dette skal da brukes til i Reslicen.
	byteslice1 := make([]byte, 55, 100)
	fmt.Println()
	// For å illustere lengden og kapasiteten til byteslice1 skriver vi inn dette:
	fmt.Printf("len: %d, cap: %d\n", len(byteslice1), cap(byteslice1))
	fmt.Println()
	// pointer til byteslice1
	fmt.Println("&byteslice1[0]")

	// skriver ut pointeren til byteslice med plass nr 0
	fmt.Println(&byteslice1[0])

	aslice := slice.Reslice(byteslice1, 50, 100)
	fmt.Println()
	// pointer til aslice. som skriver ut fra Reslice både byteslice1, 50 (lengde) og 100 (kapasitet). I dette minnet
	// som da starter på plass nr 0 i slicen.
	fmt.Println("&aslice[0]")

	fmt.Println(&aslice[0])
	fmt.Println()
	fmt.Println("&byteslice1[50]")

	// Ved bruk av pointeren blir resultatet skrevet ut slik: 0xc000048098
	fmt.Println(&byteslice1[50])
	// %byteslice1 syntax gir minneadressen av byteslice1, for eksempel en pointer til bs1.
	// Dette er fordi den har en referanse til memory adressen for den variabelen
}
