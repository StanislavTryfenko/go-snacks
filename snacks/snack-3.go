package snacks

/* Esercizio:

Scrivi una funzione che prende come input una lista di interi e restituisce la lunghezza della sottolista più lunga di numeri consecutivi.

Esempio:

Input: [1, 2, 3, 5, 6, 7, 9, 10, 11, 12] Output: 4 (la sottolista più lunga di numeri consecutivi è [5, 6, 7, 9])

Consigli:

Utilizza una variabile per tenere traccia della lunghezza della sottolista più lunga trovata finora.
Utilizza una variabile per tenere traccia della lunghezza della sottolista attuale.
Utilizza un loop per iterare sulla lista di input.
Utilizza la funzione range per iterare sulla lista di input.
Utilizza la condizione if per controllare se il numero attuale è consecutivo al numero precedente.

Nota: non dimenticare di gestire il caso in cui la lista di input è vuota!

Nota anche che la sottolista più lunga di numeri consecutivi può essere formata da più di due numeri consecutivi.
Ad esempio, nella lista [1, 2, 3, 4, 5], la sottolista più lunga di numeri consecutivi è [1, 2, 3, 4, 5], che ha una lunghezza di 5.
*/
func Snack3(nums []int) int {
	// Il tuo codice qui
	if nums == nil {
		return 0
	}

	var longestSublist int
	var actualSublist int

	for key, value := range nums {
		if key != 0 && value == nums[key-1]+1 {
			actualSublist++
		} else {
			// uno perché stai effettivamente guardando un numero
			actualSublist = 1
		}

		if actualSublist > longestSublist {
			longestSublist = actualSublist
		}
	}
	return longestSublist
}
