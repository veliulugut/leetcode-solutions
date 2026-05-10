package main 

import "fmt"

/*
Problem Summary

We have 3 dice.
Each die has 6 faces.

The j-th face of the i-th die contains A[i][j].

Each face has equal probability:

:contentReference[oaicite:0]{index=0}

We roll all 3 dice simultaneously.

Goal:
Find the probability that the rolled values are
exactly one 4, one 5, and one 6.

Examples of valid outcomes:
4 5 6
4 6 5
5 4 6
5 6 4
6 4 5
6 5 4

Approach

- Read 3 arrays (dice faces)
- Try all possible outcomes using 3 nested loops
- Total possible outcomes:

:contentReference[oaicite:1]{index=1}

- Count how many outcomes contain exactly:
  one 4, one 5, one 6

Final probability:

:contentReference[oaicite:2]{index=2}
	Link : https://atcoder.jp/contests/abc456/tasks/abc456_b
*/


func main() {
	var dice1 [6]int
	var dice2 [6]int
	var dice3 [6]int

	// input al
	for i := 0; i < 6; i++ {
		fmt.Scan(&dice1[i])
	}

	for i := 0; i < 6; i++ {
		fmt.Scan(&dice2[i])
	}

	for i := 0; i < 6; i++ {
		fmt.Scan(&dice3[i])
	}

	count := 0

	// bütün olasılıkları dene
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			for k := 0; k < 6; k++ {

				a := dice1[i]
				b := dice2[j]
				c := dice3[k]

				// tam olarak 4,5,6 mı kontrol et
				if (a == 4 && b == 5 && c == 6) ||
					(a == 4 && b == 6 && c == 5) ||
					(a == 5 && b == 4 && c == 6) ||
					(a == 5 && b == 6 && c == 4) ||
					(a == 6 && b == 4 && c == 5) ||
					(a == 6 && b == 5 && c == 4) {

					count++
				}
			}
		}
	}

	// toplam durum 216
	answer := float64(count) / 216.0

	fmt.Println(answer)
}
