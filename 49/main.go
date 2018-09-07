package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	tmp := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
		43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 107}

	mark := make(map[int][]string)
	for _, str := range strs {
		value := 1
		for _, v := range str {
			value *= tmp[uint8(v)-'a']
		}

		ss, ok := mark[value]
		if !ok {
			ss = make([]string, 0)
		}
		ss = append(ss, str)
		mark[value] = ss
	}
	result := make([][]string, 0)
	for _, v := range mark {
		result = append(result, v)
	}
	return result
}

func main() {
	ss := groupAnagrams([]string{"ron", "huh", "gay", "tow", "moe", "tie", "who", "ion", "rep", "bob", "gte", "lee", "jay", "may",
		"wyo", "bay", "woe", "lip", "tit", "apt", "doe", "hot", "dis", "fop", "low", "bop", "apt", "dun", "ben", "paw", "ere", "bad",
		"ill", "fla", "mop", "tut", "sol", "peg", "pop", "les"})
	for i := 0; i < len(ss); i++ {
		fmt.Println(ss[i])
	}
}
