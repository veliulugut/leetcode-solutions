package main


import "fmt"

// Given an unsorted integer array,
// find all pairs whose sum equals the target value.
//
// Example:
// nums = [8, 7, 2, 5, 3, 1]
// target = 10
//
// Output:
// (8,2)
// (7,3) 
// O(n^2)
func findPair(nums []int,target int){
	n := len(nums)-1

	found := false

	for i:=0;i<n-1;i++{
		for j:=i+1;j<n;j++{
		
			if nums[i] + nums[j] == target{
				
				fmt.Printf("Pair found (%v,%v)\n",nums[i],nums[j])
				
				found = true	
			}
		}
	}
		if !found{
					fmt.Println("Pair not found")
		}

}

func main(){

	nums := []int{8,7,2,5,3,1}

	findPair(nums,10)
}
