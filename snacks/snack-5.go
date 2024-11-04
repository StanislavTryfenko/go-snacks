package snacks

/* Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.

The following rules define a valid string:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "(*)"
Output: true
Example 3:

Input: s = "(*))"
Output: true


Constraints:

1 <= s.length <= 100
s[i] is '(', ')' or '*'. */

func Snack5(s string) bool {

	minOpenParentheses := 0
	maxOpenParentheses := 0
	//Per sapere se la stringa sia valida o meno a noi basta contare le parentesi aperte, un pó  stile blackjack con un mazzo di carte
	for _, c := range s {

		// Per ogni parentesi aperta aumentiamo entrambi i contatori
		if c == '(' {
			minOpenParentheses++
			maxOpenParentheses++
		}

		// Per ogni parentesi chiusa diminuiamo entrambi i contatori
		if c == ')' {
			minOpenParentheses--
			maxOpenParentheses--
		}

		// E per ogni asterisco siamo flessibili;
		// se fosse una parentesi chiusa potremmo aver bisogno di una aperta in meno
		// se fosse una parentesi aperta potremmo aver bisogno di una chiusa in meno
		if c == '*' {
			minOpenParentheses--
			maxOpenParentheses++
		}

		// Se andiamo sotto lo 0 con le maxi significa che la stringa é per forza falsa perché una parentesi chiusa si trova prima di una parentesi aperta
		if maxOpenParentheses < 0 {
			return false
		}
	}

	// Se alla fine siamo a 0 significa che siamo in pari con le parentesi, per ogni numero sotto lo 0 abbiamo avuto un asterisco "jolly"
	return minOpenParentheses <= 0
}
