package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 0 {
		return ""
	}
	str := "1"
	for i := 1; i < n; i++ {
		tmp := str[0]
		cnt := 0
		tmpstr := ""
		for j := 0; j < len(str); j++ {
			if str[j] != tmp {
				tmpstr += string(('0' + cnt)) + string(str[j-1])
				cnt = 0
				tmp = str[j]
			}
			cnt++
		}
		tmpstr += string(('0' + cnt)) + string(str[len(str)-1])
		str = tmpstr
	}
	return str
}

/*
1.     1
2.     11
3.     21
4.     1211
5.     111221
6.	   312211
7.	   13112221
8.	   1113213211
9.	   31131211131231
*/

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(countAndSay(i))
	}
	//fmt.Println(countAndSay(4))
}
