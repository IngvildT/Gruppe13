package tcpserver

//Basert på koden fra Oppgaveteksten ICA03 oppgave 5 hentet fra:
//https://uia.instructure.com/courses/3171/assignments/12048?module_item_id=85316

import (
	"net"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA3/Oppgave5/json"
)

/** OPPGAVE 5 A og B.
"func handler" funksjonen henter JSON strukturen fra json.go pakken
og skriver til klienter ved forespørsel. */
func handler(c net.Conn) {
	json := json.JSON()
	c.Write(json)
	c.Close()
}

/** TCPServer starter serveren og søker ("listens") etter klienter. "Handler" funksjonen er en go-routine
som kalles for hver tilkobling som blir akseptert */
func TCPServer() {
	l, err := net.Listen("tcp", ":5002")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
