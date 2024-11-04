package snacks

import (
	"strings"
)

/*
	Esercizio:

Scrivi una funzione che prende come input una stringa e restituisce la stringa con tutte le vocali eliminate.

Esempio:

Input: "Hello World" Output: "Hll Wrld"

Consigli:

Utilizza la funzione strings.Replace() per sostituire le vocali con una stringa vuota.
Utilizza un loop per iterare sulla stringa di input e controllare ogni carattere.
Utilizza la funzione strings.ToLower() per convertire la stringa di input in minuscolo, in modo da poter controllare facilmente le vocali.
*/
func Snack1(s string) string {
	// Il tuo codice qui

	sLower := strings.ToLower(s)

	sLowerArray := strings.Split(sLower, "")

	for i := 0; i < len(sLowerArray); i++ {
		if sLowerArray[i] == "a" || sLowerArray[i] == "e" || sLowerArray[i] == "i" || sLowerArray[i] == "o" || sLowerArray[i] == "u" {
			sLowerArray[i] = ""
		}
	}

	return strings.Join(sLowerArray, "")
}
