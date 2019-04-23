package iso

import (
	"fmt"
	"testing"
)

var testGreetingExtendedASCII1 = []struct {
	e1       string
	expected string
}{
	{"Salut, ça va °-", ""},
	{"Salut, ça va °-) Κοστίζει ​€50 Forstår du?​",""},
}

//OPPGAVE 2B
//Tester om returverdien(av type string) i funksjonen GreetingExtendedASCII1 kun inneholdeer tegn 
//fra Extended ASCII.
//Rangen for utvidet ASCII utifra det hexadesimale systemet er 7F- FF. 
func TestRangeGreetingExtendedASCII(t *testing.T) {
s := GreetingExtendedASCII1()
	for _, i := range s {
		//Dersom symbolene (i) i structen er vanlig ASCII(7 bits) altså alle verdiene er representert 
		//til og med opptil 7E i hexadesimal, og/eller dersom symbolene har en annen representasjon enn 
		//rangen mellom 7F til FF vil testen faile. 
		if i < 0x7F && i > 0xFF {
			t.Errorf("testGreetingExtendedASCII1() returns non-extended ASCII value: %q - %v", i, i)
			fmt.Printf("%c", i)
		}
		fmt.Printf("%c", i)
	}

}
