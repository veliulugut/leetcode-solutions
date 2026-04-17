package main

import(
  "fmt"
)
/*
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
*/

func isPolid(n int) bool {
	reverse := 0
	temp := n

	if n < 0 {
	return false 
	}

	for temp != 0 {
	dig := temp % 10
	reverse = reverse * 10 + dig
	temp = temp / 10

	}
	
	return reverse == n
}	




func main() {
	n := 12321

	if isPolid(n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}



