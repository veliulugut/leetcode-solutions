package main

import "fmt"


const MOD = 998244353


func solve(s string)int{
	n := len(s)

	// her tek karakter zaten geçerli
	answer := 1

	// current = şu anki geçerli zincirin uzunluğu
	current := 1

	for i:=1;i<n;i++{

		// yan yana farklıysa zincir devam eder
		if s[i] != s[i-1]{
			current ++
		}	else{
			//aynıysa sıfırdan basla
			current = 1
		}

		answer += current
		answer %=MOD
	}

	return answer
}

func main(){
	var s string
	fmt.Scan(&s)

	fmt.Println(solve(s))
}
