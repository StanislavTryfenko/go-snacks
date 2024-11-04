package snacks

/*
	Esercizio:

Scrivi una funzione che prende come input una mappa (map) di stringhe a interi e restituisce la stringa con il valore intero più alto. Se ci sono più stringhe con lo stesso valore intero più alto, restituisci tutte le stringhe.

Esempio:

Input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 3} Output: []string{"c", "d"}

Consigli:

Utilizza la funzione range per iterare sulla mappa di input.
Utilizza una variabile per tenere traccia del valore intero più alto trovato finora.
Utilizza una slice per memorizzare tutte le stringhe con il valore intero più alto.
Utilizza la funzione append per aggiungere elementi alla slice.

Nota: non dimenticare di gestire il caso in cui la mappa di input è vuota!
*/
func Snack2(m map[string]int) []string {
	// Il tuo codice qui
	if m == nil {
		return []string{}
	}

	var maxValue int
	var keys []string

	for key, value := range m {
		if value > maxValue {
			maxValue = value
			keys = []string{key}
		} else if value == maxValue {
			keys = append(keys, key)
		}
	}
	return keys
}
