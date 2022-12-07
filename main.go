package main

import (
	"fmt"
)

/**
 *
 * Write a program that prints out the numbers 1 to 100 (inclusive).
 *
 * If the number is divisible by 3, print Crackle instead of the number.
 * If it's divisible by 5, print Pop.
 * If it's divisible by both 3 and 5, print CracklePop.
 * */
func main() {
	CracklePop(100)
}

func CracklePop(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(GetOutput(i))
	}
}

func GetOutput(i int) string {
	divisibleByThree := i%3 == 0
	divisibleByFive := i%5 == 0

	if divisibleByThree && divisibleByFive {
		return "CracklePop"
	}

	if divisibleByThree {
		return "Crackle"
	}

	if divisibleByFive {
		return "Pop"
	}

	return fmt.Sprintf("%d", i)
}
