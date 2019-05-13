package bfrequence

import (
    "bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

/*
Frequence bruker flere funksjoner, og til slutt et map for
å printe ut en oversikt over de mest brukte runene i en valgfri fil.

Vi har også funksjonen sjekkLinjer som returnerer antall linjer i en fil.

*/

//Sjekker antall linjer og returnerer disse som en int.
func SjekkLinjer(filInfo string, finn string) int {
	file, err := os.Open(filInfo)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Leser av input", err)
	}
	fmt.Printf("Antall linjer i filen er: %d\n", count)
	return count
}

//Sjekker forekomsten av runer.
func SjekkRuner(filinfo string, finn string) int {
	fil, err := os.Open(filinfo)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 512*1024)
	counter := 0
	searchFor := []byte(finn)

	c, _ := fil.Read(buffer)
	counter += bytes.Count(buffer[:c], searchFor)
	return counter
}

//Returnerer et map.
func SjekkForekomst(filInfo string) map[int]string {
	m := make(map[int]string)

	for i := 0x20; i <= 0x7F; i++ {
		count := SjekkRuner(filInfo, string(i))
		rune := string(i)
		m[count] = rune
	}
	fmt.Println()
	return m
} 

//Tar inn et map og printer ut de mest vanlige runene.
func RuneFrekvens(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Printer de fem mest brukte runene
	fmt.Println("Her er en oversikt over de mest vanlige runene.")
	verdi1 := keys[len(keys)-1]
	verdi2 := keys[len(keys)-2]
	verdi3 := keys[len(keys)-3]
	verdi4 := keys[len(keys)-4]
	verdi5 := keys[len(keys)-5]
	fmt.Println("1. Rune: ", m[verdi1], "Antall: ", verdi1)
	fmt.Println("2. Rune: ", m[verdi2], "Antall: ", verdi2)
	fmt.Println("3. Rune: ", m[verdi3], "Antall: ", verdi3)
	fmt.Println("4. Rune: ", m[verdi4], "Antall: ", verdi4)
	fmt.Println("5. Rune: ", m[verdi5], "Antall: ", verdi5)
}