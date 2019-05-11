package json

//OPPGAVE 5B
//https://golang.org/pkg/encoding/json/#example_Marshal

import (
	"encoding/json"
)

//Dette er struct(samling av feltene) for Navn og Email.
type NameEmail struct {
	FirstName string
	LastName  string
	Email     string
}

/** Feltene (structen) defineres deretter i "func JSON" hvor den blir gjort om til JSON struktur
og strukturen blir deretter returnert til klienten via "var b" */
func JSON() []byte {
	nameEmail := NameEmail{
		FirstName: "Ingvild",
		LastName:  "Tisland",
		Email:     "ingvild.tisland@gmail.com",
	}

	var b []byte
	b, err := json.Marshal(nameEmail)
	if err != nil {
		panic(err)
	}
	return b
}
