package main

import (
	tcpserver "Gruppe13/ICA3/Oppgave5/tcp_server"
	//tcpclient "Gruppe13/ICA3/Oppgave5/tcp_client"
)

/** Kjører TCPServer og dersom det kommer forespørsel fra tilkoblet klient returneres
JSON strukturen til klienten via tcp_server.go */
func main() {
	//tcpclient.TCPClient()
	//tcpclient.TCPClient2()
	tcpserver.TCPServer()

}
