package main

import (
	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave3/bfrequence"
	//"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave3/frequence"
)

func main() {
	const file = "./pg100.txt"
	bfrequence.SjekkLinjer(file, "\n")
	b := bfrequence.SjekkForekomst(file)
	bfrequence.RuneFrekvens(b)

	//frequence.SjekkAntallLinjer(file)
	//f := frequence.FåØversteRuner(file)
	//frequence.TellRuner(f)
}
