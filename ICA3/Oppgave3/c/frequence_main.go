package main

import (
	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave3/bfrequence"
	//"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave3/frequence"
)

func main() {
	const file = "./pg100.txt"
	
	//bfrequence
	bfrequence.SjekkLinjer(file, "\n")
	b := bfrequence.SjekkForekomst(file)
	bfrequence.SjekkRuner(file, "pg100.txt")
	bfrequence.RuneFrekvens(b)


	//frequence
	//frequence.SjekkAntallLinjer(file)
	//frequence.TellRuner(file)
	//frequence.FåØversteRuner(file)
	//frequence.AntallBytes(file)
}
