package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

)

func main() {
	var filnavn string

	flag.StringVar(&filnavn, "f", "", "filename")

	flag.Parse()

	bufioFindLines(filnavn)
}

func bufioFindLines(filnavn string) {
	file, err := os.Open(filnavn)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("number of lines:", lineCount)
}