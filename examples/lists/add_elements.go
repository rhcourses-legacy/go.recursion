package lists

// AddElements erwartet eine Liste von Zahlen.
// Die Funktion liefert die Summe aller Elemente dieser Liste.
func AddElements(list []int) int {
	// Basisfall: Wenn die Liste leer ist, ist die Summe 0.
	if len(list) == 0 {
		return 0
	}

	// Addiere das erste Element zur Summer aller weiteren Elemente.
	return list[0] + AddElements(list[1:])
}
