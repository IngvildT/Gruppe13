package main

//Kode laget utifra eksempel på: https://dlintw.github.io/gobyexample/public/http-client.html

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/** OPPGAVE 5C. "func main" sender forespørsel til serveren :8989 om å få koble seg til.
http.Get metoden returnerer 2 verdier. 1: respons objektet som i dette tilfellet er "ok"
og 2: Hvis det skjer en feil i tilkoblingen til serveren og å få en respons tilbake, altså
hvis err ikke er lik nil vil det komme en feilmelding  */

func main() {
	resp, err := http.Get("http://localhost:5000/ok")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//Når klienten har fått kontakt med serveren leser den "bodyen fra serveren"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//Skriver "bodyen" fra serveren som i dette tilfellet er nameEmail strukturen hentet fra json.go
	_, err = os.Stdout.Write(body)
	if err != nil {
		log.Fatal(err)
	}

}
