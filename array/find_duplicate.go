package main

import "fmt"

/*
Input:  { 1, 2, 3, 4, 4 }
Output: The duplicate element is 4 
Input:  { 1, 2, 3, 4, 2 }
Output: The duplicate element is 2
*/

func main() {
	arr := []int{ 1, 2, 3, 4, 4 }

		//findDuplicates(arr)
		 
		optimizeFindDup(arr)
	}

// O(n²)
func findDuplicates(data []int) {
	 var (
		 dublicates []int 
			 )

	 for i := 0 ; i < len(data) ; i ++ {
		  for j := i+ 1 ; j < len(data) ; j++ {
				 if data[i] == data[j] {
					 dublicates = append(dublicates,data[i])
				 } 
			}
	 }

	fmt.Println(dublicates) 
}

// hashmap O(n)
func optimizeFindDup(data []int) {
	values := make(map[int]struct{})

	var dublicates []int

	for _ , d := range data {
		if _ , ok := values[d]; ok {
			dublicates = append(dublicates,d)
			continue
		}

		values[d] = struct{}{}
	}

	fmt.Println(dublicates)
}




