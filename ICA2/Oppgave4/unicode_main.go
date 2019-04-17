package main

import (
	"fmt"
	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA2/Oppgave4/unicode"
)

func main()  {
	Is := unicode.Translate("nord og sør", "Is")
	Jp := unicode.Translate("nord og sør", "Jp")
		fmt.Println(Is)
		fmt.Println(Jp)

	unicode.UnicodeCodePointDemo()	
}

