package main

import (
	"fmt"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA1-02/Oppgave3/sum"
)

func main() {
	finish := false
	for finish == false {

		var typeTall string
		fmt.Println("Gyldige typer: uint32, int32, int64, float64.")
		fmt.Println("Skriv inn hvilke type du ønsker å regne med")
		fmt.Scan(&typeTall)
		regneType := string(typeTall)
		fmt.Println("Du har valgt å skrive med " + typeTall)

		if regneType == "uint32" {
			var typen uint32
			typen2 := 0
			fmt.Println("Skriv inn det første tallet. ")
			fmt.Scan(&typen)
			tallet1 := uint32(typen)
			fmt.Println("Skriv inn det andre tallet. ")
			fmt.Scan(&typen2)
			tallet2 := uint32(typen2)
			fmt.Println("\nSummen er: ", sum.SumUint32(tallet1, tallet2))
		} else if regneType == "int32" {
			var typen int32
			typen2 := 0
			fmt.Println("Skriv inn det første tallet.")
			fmt.Scan(&typen)
			tallet1 := int32(typen)
			fmt.Println("Skriv inn det andre tallet.")
			fmt.Scan(&typen2)
			tallet2 := int32(typen2)
			fmt.Println("\nSummen er: ", sum.SumInt32(tallet1, tallet2))
		} else if regneType == "int64" {
			var typen int64
			typen2 := 0
			fmt.Println("Skriv inn det første tallet.")
			fmt.Scan(&typen)
			tallet1 := int64(typen)
			fmt.Println("Skriv inn det andre tallet.")
			fmt.Scan(&typen2)
			tallet2 := int64(typen2)
			fmt.Println("\nSummen er: ", sum.SumInt64(tallet1, tallet2))
		} else if regneType == "float64" {
			var typen float64
			typen2 := 0
			fmt.Println("Skriv inn det første tallet.")
			fmt.Scan(&typen)
			tallet1 := float64(typen)
			fmt.Println("Skriv inn det andre tallet.")
			fmt.Scan(&typen2)
			tallet2 := float64(typen2)
			fmt.Println("\nSummen er: ", sum.SumFloat64(tallet1, tallet2))

		} else {
			fmt.Println("Dette er dessverre ikke en gyldig verdi. Prøv igjen.")
		}

		var avslutt string

		fmt.Println("Avslutt ved å skrive avslutt. Fortsett med å skrive en random bokstav")
		fmt.Scan(&avslutt)

		if avslutt == "avslutt" {
			finish = true

		} else {
			finish = false
		}

	}
}
