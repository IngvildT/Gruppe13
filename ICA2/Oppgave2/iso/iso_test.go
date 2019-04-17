package iso

import (
	"fmt"
	"testing"
)

//OPPGAVE 2B
//Test funksjon for TestGreetingASCII funksjonen fra iso.go
func TestGreetingASCII(t *testing.T) {
	if !(rangeExtendedASCII(GreetingExtendedASCII())) {
		t.Fail()
	}
}

func rangeExtendedASCII(s string) bool {
	for _, c := range s {
		fmt.Println(c)
		if c == 256 {
			return false
		}
	}
	return true
}
