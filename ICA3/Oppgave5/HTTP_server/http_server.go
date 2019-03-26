package main

import (
	"fmt"
	"net/http"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA3/Oppgave5/json"
)

func main() {
	http.HandleFunc("/ok", handler)
	http.ListenAndServe(":8989", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := json.JSON()
	fmt.Fprintf(w, string(s))
}
