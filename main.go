package main

import (
	"fmt"
	"snacks"
)

func main() {
	fmt.Println(snacks.Snack1("Abaco"))

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 3}
	fmt.Println(snacks.Snack2(m))

	nums := []int{1, 2, 3, 5, 6, 7, 9, 10, 11, 12}
	fmt.Println(snacks.Snack3(nums))

	toEncrypt1 := []int{1, 2, 3}
	toEncrypt2 := []int{10, 21, 31}
	toEncrypt3 := []int{241, 352, 231, 3333332735}

	fmt.Println(snacks.Snack4(toEncrypt1))
	fmt.Println(snacks.Snack4(toEncrypt2))
	fmt.Println(snacks.Snack4(toEncrypt3))

	testSnack5 := [5]string{
		"()",
		"(*)",
		"(*))",
		"((*)",
		"((**(",
	}
	for _, v := range testSnack5 {
		fmt.Println(snacks.Snack5(v))
	}
}
