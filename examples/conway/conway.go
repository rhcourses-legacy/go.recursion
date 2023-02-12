package conway

// Element ist ein Datentyp, der ein Element der Element-Folge darstellt.
type Element []int

// Next liefert zu einem aktuellen Element das nächste.
// Man beachte: Diese Funktion ist nicht rekursiv.
func (current Element) Next() Element {
	result := Element{}

	count := 1
	for i, el := range current[1:] {
		if el == current[i] {
			count++
		} else {
			result = append(result, count, current[i])
			count = 1
		}
	}
	result = append(result, count, current[len(current)-1])

	return result
}

// Conway berechnet alle Elemente der Conway-Folge mit einem bestimmten Startwert
// und bis zum gegebenen n. Diese Funktion ist nun rekursiv.
func Conway(start, n int) []Element {
	if n == 0 {
		return []Element{{start}}
	}
	elements := Conway(start, n-1)
	return append(elements, elements[len(elements)-1].Next())
}
