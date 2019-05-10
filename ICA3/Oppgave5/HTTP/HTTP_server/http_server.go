package main

import (
	"fmt"
	"net/http"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave5/json"
)
/** OPPGAVE 5C 
func main starter serveren og søker etter klienter. */

func main() {
	http.HandleFunc("/ok", handler)
	http.ListenAndServe(":5000", nil)
}

/** OPPGAVE 5C "func handler" funksjonen henter JSON strukturen fra json.go pakken
og skriver til klienter ved forespørsel. */
func handler(w http.ResponseWriter, r *http.Request) {
	s := json.JSON()
	fmt.Fprintf(w, string(s))
}
