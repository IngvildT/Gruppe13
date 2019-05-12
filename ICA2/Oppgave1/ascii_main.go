package main

import (
	"fmt"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA2/Oppgave1/ascii"
)

const Ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

//Kjører ascii.go
func main() {

	ascii.IterateOverASCIIStringLiteral(ascii.Ascii)
	fmt.Printf(ascii.GreetingASCII())

}
