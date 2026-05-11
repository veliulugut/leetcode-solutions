package main 

import "fmt"

/*
Input:  { 1, 0, 1, 0, 1, 0, 0, 1 }
Output: { 0, 0, 0, 0, 1, 1, 1, 1 }
*/ 


func binearySort(arr []int)[]int{
    count := 0 

		for i := 0 ; i < len(arr) ; i++{
			 if arr[i] == 0 {
				 count ++
			 }
		}

		for j := 0 ; j < len(arr) ; j++ {
				if j < count {
					arr[j] = 0
				} else {
					arr[j] = 1
				}
		}

		return arr

}

func main() {
	x := []int{ 1, 0, 1, 0, 1, 0, 0, 1 }
	
	fmt.Println(binearySort(x))

}

