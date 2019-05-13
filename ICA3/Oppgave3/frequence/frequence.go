package frequence

import (
    "bufio"
    "io"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

//SjekkAntallLinjer returnerer hvor mange linjer filen har som int
func SjekkAntallLinjer(filNavn string) int {

	filInnhold, err := os.Open(filNavn)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(filInnhold)

	linjer := 0
	for fileScanner.Scan() {
		linjer++
	}
	return linjer
}

//TellRuner bruker bufio sin NewReader, mer spesifikt ReadRune for å populere et map av hvilke runes som gjentas mest.
func TellRuner(filNavn string) map[rune]int {

	teller, err := os.Open(filNavn)
	if err != nil {
		log.Fatal(err)
	}
	leser := bufio.NewReader(teller)
	runer := map[rune]int{}

	for {
		r, _, err := leser.ReadRune()
		if err == io.EOF {
			break
		}

		runer[r]++
	}

	return runer
}

func FåØversteRuner(filNavn string) {
	teller := TellRuner(filNavn)

	var nøkkel []int

	for _, v := range teller {
		nøkkel = append(nøkkel, v)

	}
	sort.Ints(nøkkel)

}

//AntallBytes returnerer et byte-array.
func AntallBytes (filNavn string) []byte {
	filinnhold, err := ioutil.ReadFile(filNavn)
	if err != nil {
		log.Fatal(err)
	}
	return filinnhold
}