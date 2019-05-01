package main

import (
	"fmt"

	slice "Gruppe13/ICA2/Oppgave5/slice/mainS"
)

func main() {

	// Printer ut funksjonene hvor vi allokerer og deklarere slice
	fmt.Println(slice.AllocateMake(6))
	fmt.Println(slice.AllocateVar(4))

	// Oppretter byteslice1
	// bruker make til å lage en slice og setter inn byte, lidx int, uidx int.
	// dette skal da brukes til i Reslicen
	byteslice1 := make([]byte, 51, 100)

	// pointer til byteslice1
	fmt.Println("&byteslice1[0]")

	fmt.Println(&byteslice1[0])

	aslice := slice.Reslice(byteslice1, 50, 100)

	fmt.Println("&aslice[0]")

	fmt.Println(&aslice[0])

	fmt.Println("&byteslice1[50]")

	// Ved bruk av pointeren blir resultatet skrevet ut slik: 0xc000048098
	fmt.Println(&byteslice1[50])
	// %byteslice1 syntax gir minneadressen av byteslice1, for eksempel en pointer til bs1.
	// Dette er fordi den har en referanse til memory adressen for den variabelen
}
