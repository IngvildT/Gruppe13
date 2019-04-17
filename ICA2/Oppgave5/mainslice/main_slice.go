package main

import (
	"fmt"
)

//Ved å ta i bruk make, kan vi lage en empty slice, uten å måtte
// lage en array

func main() {
	fmt.Println(AllocateVar(5))

	fmt.Println(AllocateVar(7))

	fmt.Println(AllocateMake(2), "&byteslice1[0]")

}
func AllocateMake(b int) []byte {

	arr := make([]byte, b)
	return arr
}
func AllocateVar(b int) []byte {

	var arr []byte = make([]byte, b)
	return arr
}

//func Reslice(slc []byte, lidx int, uidx int) []byte {

//Funksjonen Resclice skal bruke en av Allocate... funksjonene for å allokere plass i minnet for en “slice”.
//Sjekk ut hva følgende kode viser (dette er implementer i en av main-funksjonene, som da er lokert en mappe “høyere” i filsystemets hierarkiet:

//}
