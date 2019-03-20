package main

import (
	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA3/Oppgave5/tcp_server"
	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA3/Oppgave5/tcp_client"
)

/** Kjører TCPServer og dersom det kommer forespørsel fra tilkoblet klient returneres 
JSON strukturen til klienten via tcp_server.go */
func main() {
	go tcpclient.TCPClient()
	tcpserver.TCPServer()

}