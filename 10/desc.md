10. Regular Expression Matching

Implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true


解题思路， 这题可以用递归的方法来解决， 但是使用动态规划更佳,自底向上的解决问题

用 dp[i][j] 来表示 s[0:i], p[0:i] 是否匹配


分解成子问题
* p[i-1] != '*'
   * if p[j-1] = '.'  then dp[i][j] = dp[i-1][j-1]
   * s[i-1] == p[j-1]  then dp[i][j] = dp[i-1][j-1]
   * then s[i][j] = false
* p[i-1] == '*'
    * if p[i-2] == s[i-1]  then dp[i][j]=dp[i-1][j]
    * if p[i-2] == '.' then dp[i][j] = dp[i-1][j]
    * then p[i][j] = dp[i-2][j]
    

分解成来子问题来之后可以按照自顶向下或自底向上的方法来求解问题

