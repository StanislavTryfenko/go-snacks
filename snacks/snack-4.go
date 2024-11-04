package snacks

import (
	"strconv"
	"strings"
)

/*
	You are given an integer array nums containing positive integers.
	We define a function encrypt such that encrypt(x) replaces every digit in x with the largest digit in x. For example, encrypt(523) = 555 and encrypt(213) = 333.

Return the sum of encrypted elements.

Example 1:

Input: nums = [1,2,3]

Output: 6

Explanation: The encrypted elements are [1,2,3]. The sum of encrypted elements is 1 + 2 + 3 == 6.

Example 2:

Input: nums = [10,21,31]

Output: 66

Explanation: The encrypted elements are [11,22,33]. The sum of encrypted elements is 11 + 22 + 33 == 66.

Constraints:

1 <= nums.length <= 50
1 <= nums[i] <= 1000
*/
func Snack4(nums []int) int {

	var result int
	var encryptedNums []int

	for i := 0; i < len(nums); i++ {
		str := strconv.Itoa(nums[i])
		var figures []int
		var higtestFigure int
		var encryptedArrStr []string

		for ii := 0; ii < len(str); ii++ {
			figures = append(figures, int(str[ii]-'0'))
		}

		for ii := 0; ii < len(figures); ii++ {
			if figures[ii] > higtestFigure {
				higtestFigure = figures[ii]
			}
		}

		for ii := 0; ii < len(str); ii++ {
			encryptedArrStr = append(encryptedArrStr, strconv.Itoa(higtestFigure))
		}
		encryptedStr := strings.Join(encryptedArrStr, "")
		encryptedNum, _ := strconv.Atoi(encryptedStr) // l'errore sicuramente non mi servirá
		encryptedNums = append(encryptedNums, encryptedNum)
	}

	for i := 0; i < len(encryptedNums); i++ {
		result += encryptedNums[i]
	}

	return result
}
