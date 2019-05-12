package main

import (
	"fmt"

	"github.com/IngvildT/IS105-Gruppe13-mappe/ICA-REPO/ICA1-02/Oppgave5/boring"
)

// Kode fra Rob Pike https://talks.golang.org/2012/concurrency.slide#20
func main() {
	c := make(chan string)
	go boring.Boring10("boring!", c)
	for i := 0; ; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
	}
	fmt.Println("You're boring; I'm leaving.")

}

