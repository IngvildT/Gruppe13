package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filnavn := os.Args[1]
	fileInfo(filnavn)
}

func fileInfo(filnavn string) {
	fileInfo, err := os.Lstat(filnavn)
	if err != nil {
		log.Fatal(err)
	}

	i64 := fileInfo.Size()
	bytes := float64(i64)
	KiB := bytes / 1024
	MiB := KiB / 1024
	GiB := MiB / 1024

	fmt.Println("Size: \n", bytes, "bytes \n", KiB, "kibibytes \n", MiB, "mibibytes \n", GiB, "gigibytes")

	// sjekker om filen er en mappe
	if fileInfo.IsDir() == true {
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}

	// sjekker om filen er en regular file
	if fileInfo.Mode().IsRegular() == true {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}

	// Has Unix permission bits: -rwxrwxrwx
	fmt.Println("Has Unix permission bits:", fileInfo.Mode().Perm())

	// Is/Is not append only (virker ikke)
	//if fileInfo.Mode().ModeAppend() == true {
	//	fmt.Println("Is append only")
	//} else {
	//	fmt.Println("Is not append only")
	//}

	// Is/is not a device file

	// Is/is not a unix character device

	//Is/is not a unix block device

	// Is/is not a symbolic link

}
