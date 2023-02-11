package ackermann

// Ack ist die "einfache" Variante der Ackermannfunktion.
// Damit ist gemeint, dass die Ackermannfunktion nur mit Rekursion,
// aber ohne jegliches Caching berechnet wird.
func Ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return Ack(m-1, 1)
	}
	return Ack(m-1, Ack(m, n-1))
}
