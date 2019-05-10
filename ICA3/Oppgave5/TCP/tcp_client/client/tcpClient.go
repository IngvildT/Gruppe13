package tcpclient

//https://golang.org/pkg/net/

import (
	"fmt"
	"io/ioutil"
	"net"
)

/** OPPGAVE 5A
TCPClient funksjonen sender forespørsel til server om tilkobling */

func TCPClient() {
	fmt.Println("Søker etter server")
	conn, err := net.Dial("tcp", "172.20.10.4:5003")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}

func TCPClient2() {
	fmt.Println("Søker etter server")
	conn, err := net.Dial("tcp", "172.20.10.5:5003")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
