package main

// Hentet fra https://www.devdungeon.com/content/working-files-go
import (
	"log"
	"os"
)

/** Denne funksjonen bryter filen ned til å bestå av 100 bytes. Er innholdet i filen mindre enn 100
bytes vil det originale innholde bestå slik det gjorde også vil resten bli fyllt med bytes tilsvarende
null. Hvis innholdet er over 100 bytes vil alt over 100 bytes gå tapt. Man kan endre parameter etter
hvor mange bytes man vil "trunctere" filen til. */

func main() {

	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
