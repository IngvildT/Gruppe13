package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var filnavn string
	flag.StringVar(&filnavn, "f", "", "filename")
	flag.Parse()

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

	// Is/Is not append only
	if fileInfo.Mode()&os.ModeAppend != 0 {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}

	// Is/is not a device file
	if fileInfo.Mode()&os.ModeDevice != 0 {
		fmt.Println("Is a device file")

		if fileInfo.Mode()&os.ModeCharDevice != 0 {
			fmt.Println("Is a Unix character device")
			fmt.Println("Is not a Unix block device")
		} else {
			fmt.Println("Is not a Unix character device")
			fmt.Println("Is a Unix block device")
		}
	} else {
		fmt.Println("Is not a device file")
		fmt.Print("Is not a Unix character device")
		fmt.Println("Is not a Unix block device")
	}
}
