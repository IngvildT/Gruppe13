package slice

func AllocateMake(b int) []byte {
	//
	s := make([]byte, b)
	// returnerer en slice av typen byte ved bruk av make metoden
	return s
}
func AllocateVar(b int) []byte {

	// s er en variabel (slice). Bruker både var og make til å lage en slice med argument b
	// og med typen byte
	var s []byte = make([]byte, b)

	// Returnerer en slice av typen byte
	return s
}

// Reslice funksjonen skal slice en allerede eksisterende slice
// ved bruk av en av Allocatefunksjonene
func Reslice(a []byte, lidx int, uidx int) []byte {

	//lager en ny slice av allocateMake til å allokere plass i minnet til en slice
	slice1 := AllocateMake(20)

	slice1 = a[lidx:uidx]

	return slice1

}
