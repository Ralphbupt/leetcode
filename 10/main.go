package main

import "fmt"

func main() {
	test("ssss", ".*", true)
	test("ssss", "ss", false)

	test("s", ".*", true)

	test("s", "s", true)

	test("s", "s", true)

	test("as", ".*", true)
	test("aab", "c*a*b", true)
	test("ab", ".*", true)

}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i, _ := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else {
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			}

		}
	}
	return dp[m][n]
}

func test(s string, p string, expected bool) {
	if isMatch(s, p) != expected {
		fmt.Printf("error while testing %s %s , expected %v, get %v", s, p, expected, isMatch(s, p))
	}
}
