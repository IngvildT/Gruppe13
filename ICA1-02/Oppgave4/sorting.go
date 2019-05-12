package algorithms

// Les https://en.wikipedia.org/wiki/Bubble_sort


func Bubble_sort_modified(list []int) {
	//Modifisert kode ved bruk av Golang´s Tuple Assignment

	n := len(list)
	swapped := true

	for swapped {
		swapped = false
		// Iterer over alle elementene i lista
		for i := 1; i < n; i++ {
			// Hvis det aktuelle elementet er større enn det neste
			// elementet, skal disse swappes
			if list[i-1] > list[i] {
				// swap elementene ved å bruke Go's Tuple Assignment
				list[i], list[i-1] = list[i-1], list[i]
				// Sett swapped til true - viktig!
				// Hvis loopen tar slutt og swapped enda er lik false vil algoritmen anta
				// at listen er sortert.
				swapped = true
			}
		}
	}
}

// Implementering av Bubble_sort algoritmen
func Bubble_sort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
