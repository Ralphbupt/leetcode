package main

import (
	"fmt"
	"reflect"
)

func findSubstring(s string, words []string) []int {
	result := make([]int, 0)
	slen, wlen := len(s), len(words)

	if slen <= 0 || wlen <= 0 {
		return result
	}
	w := len(words[0])

	for i := 0; i <= slen-wlen*w; i++ {
		dic := make(map[string]int)
		for _, key := range words {
			dic[key] += 1
		}
		for j := 0; j < wlen; j++ {
			str := s[i+j*w : i+j*w+w]
			if value, ok := dic[str]; !ok {
				break
			} else if value == 1 {
				delete(dic, str)
			} else {
				dic[str]--
			}
			if len(dic) == 0 {
				result = append(result, i)
			}
		}
	}
	return result
}

func main() {
	test("aaabbbccc", []string{"aaa", "bbb", "ccc"}, []int{0})
	test("aabbccbbaa", []string{"aa", "bb"}, []int{0, 6})
	test("aaaaaaa", []string{"a"}, []int{0, 1, 2, 3, 4, 5, 6})
	test("aaaaaaa", []string{"aa"}, []int{0, 1, 2, 3, 4, 5})
	test("safdafhusdh", []string{"ff"}, []int{})
}

func test(s string, words []string, res []int) {
	if !reflect.DeepEqual(findSubstring(s, words), res) {
		fmt.Printf("error test %s %v, expected %v, get %v", s, words, res, findSubstring(s, words))
	}
}
