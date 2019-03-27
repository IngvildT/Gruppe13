package main

//Hentet fra https://www.devdungeon.com/content/working-files-go
import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

//Denne funksjonen lager en ny tom text- fil kalt test.txt
func main() {
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}
