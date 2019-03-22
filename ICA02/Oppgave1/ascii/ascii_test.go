package ascii

import "testing"
import "fmt"

/** Oppgave 1d
Test for funksjonen ​greetingASCII() ​i egen fil ascii_test.go​,
som tester om returverdier (av type ​string​) inneholderkun ASCII-tegn. */

func TestGreetingASCII(t *testing.T) {
	if !(rangeASCII(GreetingASCII())) {
		t.Fail()
	}
}

func rangeASCII(s string) bool {
	for _, c := range s {
		fmt.Println(c)
		if c == 128 {
			return false
		}
	}
	return true
}
