package main

import (
	"fmt"
	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA2/Oppgave4/unicode"
)

func main()  {
	is := unicode.Translate("nord og sør", "is")
	jp := unicode.Translate("nord og sør", "jp")
		fmt.Println(is)
		fmt.Println(jp)

	unicode.UnicodeCodePointDemo()	
}

