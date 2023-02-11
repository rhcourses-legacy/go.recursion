package ackermann

import "fmt"

// AckCache ist eine int-Matrix, die dafür gedacht ist, einmal berechnete
// Werte der Ackermannfunktion zwischenzuspeichern, um die Berechnung zu optimieren.
type AckCache [][]int

// NewAckCache erzeugt einen neuen, leeren, AckCache.
func NewAckCache() AckCache {
	return AckCache{}
}

// AddRow fügt eine Zeile zum Cache hinzu.
// Die Zeile enthält genau so viele Einträge, wie die übrigen Zeilen auch.
// Alle Einträge sind mit 0 initialisiert.
func (cache *AckCache) AddRow() {
	*cache = append(*cache, make([]int, cache.Columns()))
}

// AddColumn fügt eine neue Spalte zum Cache hinzu.
// Die neuen Einträge sind am Anfang 0.
func (cache *AckCache) AddColumn() {
	for row := range *cache {
		(*cache)[row] = append((*cache)[row], 0)
	}
}

// ExtendTo erweitert den Cache, so dass er die Positionen m und n enthält.
func (cache *AckCache) ExtendTo(m, n int) {
	for m >= cache.Rows() {
		cache.AddRow()
	}
	for n >= cache.Columns() {
		cache.AddColumn()
	}
}

// Rows liefert die Anzahl der Zeilen des Caches.
func (cache AckCache) Rows() int {
	return len(cache)
}

// Columns liefert die Anzahl der Spalten des Caches.
func (cache AckCache) Columns() int {
	if cache.Rows() == 0 {
		return 0
	}
	return len(cache[0])
}

// HasValue liefert true, falls der Cache einen Wert für die angegebene Position hat.
func (cache AckCache) HasValue(m, n int) bool {
	return m >= 0 && m < cache.Rows() &&
		n >= 0 && n < cache.Columns() &&
		cache[m][n] > 0
}

// Get liefert den Wert für m und n.
// Falls der Cache den Wert noch nicht hat, wird er berechnet.
func (cache *AckCache) Get(m, n int) int {
	if !cache.HasValue(m, n) {
		cache.Compute(m, n)
	}
	return (*cache)[m][n]
}

// Compute berechnet den Wert für m und n und speichert
// ihn an der passenden Stelle im Cache.
func (cache *AckCache) Compute(m, n int) {
	cache.ExtendTo(m, n)
	if m == 0 {
		(*cache)[m][n] = n + 1
		return
	}
	if n == 0 {
		(*cache)[m][n] = cache.Get(m-1, 1)
		return
	}
	(*cache)[m][n] = cache.Get(m-1, cache.Get(m, n-1))
}

// String liefert eine String-Repräsentation des Caches mit Zeilennummern.
func (cache AckCache) String() string {
	result := ""
	for m, row := range cache {
		result += fmt.Sprintf("%d:", m)
		for _, value := range row {
			result += fmt.Sprintf(" %2d", value)
		}
		result += "\n"
	}
	return result
}
