package main

import (
	tcpclient "github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA3/Oppgave5/tcp_client"
)

/** Kjører TCPServer og dersom det kommer forespørsel fra tilkoblet klient returneres
JSON strukturen til klienten via tcp_server.go */
func main() {
	tcpclient.TCPClient()
	//tcpclient.TCPClient2()

}
