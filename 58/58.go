package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	lenght, tail := 0, len(s)-1
	for tail >= 0 && s[tail] == ' ' {
		tail--
	}

	for tail >= 0 && s[tail] != ' ' {
		lenght++
		tail--
	}
	return lenght
}

func main() {

	fmt.Println(lengthOfLastWord("da sdasd sdsds"))
	fmt.Println(lengthOfLastWord("sdss sdsds"))
	fmt.Println(lengthOfLastWord("aaaa"))
	fmt.Println(lengthOfLastWord("sssssss s"))
}
