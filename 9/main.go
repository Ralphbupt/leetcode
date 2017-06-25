package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	p, q := 0, x
	for q > 10 {
		p = p*10 + q%10
		q = q / 10
	}
	return p == x/10 && q == x%10
}

func main() {

}
