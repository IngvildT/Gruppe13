package main

//Eksempel kode for client hentet fra: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

import "net"
import "fmt"
import "bufio"
import "os"

//Oppgave 5A
func main() {

	//Koble til socket
	conn, _ := net.Dial("tcp", "127.0.0.1:5000")
	for {
		// Leser Input fra Stdin (Standard Inputs)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Sender til socket
		fmt.Fprintf(conn, text+"\n")
		// Gir en respons fra server
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
