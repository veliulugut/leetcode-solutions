package main

import "fmt"
/*
There are three dice, each with six faces.
Each die has 
1,2,3,4,5,6 written on its faces.
When these dice are rolled simultaneously, is it possible for the sum of the rolled values to be 
X?
*/

func main () {
	var ( 
		x int
	)

	fmt.Scan(&x)

	if x >=3 && x <= 18{
			fmt.Println("Yes")
	}else{
			fmt.Println("No")
	} 
}
