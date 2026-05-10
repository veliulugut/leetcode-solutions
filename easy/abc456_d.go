package main

import "fmt"

const MOD = 998244353

func solve(s string) int {
	// a,b,c ile biten geçerli subsequence sayıları
	endA := 0
	endB := 0
	endC := 0

	for _, ch := range s {

		if ch == 'a' {

			// yeni "a"
			// veya b/c ile bitenlere a ekle
			newA := (1 + endB + endC) % MOD
			endA = (endA + newA) % MOD

		} else if ch == 'b' {

			newB := (1 + endA + endC) % MOD
			endB = (endB + newB) % MOD

		} else {

			newC := (1 + endA + endB) % MOD
			endC = (endC + newC) % MOD
		}
	}

	answer := (endA + endB + endC) % MOD

	return answer
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(solve(s))
}
