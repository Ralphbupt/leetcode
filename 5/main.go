package main

import "fmt"

func longestPalindrome(s string) string {

	slen := len(s)
	if slen <= 1 {
		return s
	}
	length := 0
	start_index := 1
	for i := 1; i < slen; i++ {
		l := i - 1
		h := i + 1
		tlen := 1
		for l >= 0 && h < slen {
			if s[l] == s[h] {
				tlen += 2
				l--
				h++
			} else {
				break
			}
		}

		if tlen > length {
			length = tlen
			start_index = l + 1
		}
		l = i - 1
		h = i
		tlen = 0
		for l >= 0 && h < slen {
			if s[l] == s[h] {
				tlen += 2
				l--
				h++
			} else {
				break
			}
		}

		if tlen > length {
			length = tlen
			start_index = l + 1
		}

	}
	return s[start_index : start_index+length]
}

func main() {
	fmt.Println(longestPalindrome("cffbbbbffc"))
}
