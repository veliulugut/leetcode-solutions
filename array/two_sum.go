package main 

/*
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/ 
import "fmt"

// Brute Force
func twoSum(arr []int ,target int)[]int{
	 for i := 0 ; i< len(arr) -1 ; i++{
		 for j:= i +1 ; j<len(arr); j++{
			 if arr[i] + arr[j] == target{
				 return []int{i,j}
			 }
		 }
	 }

	 return []int{}
}

func main () {
	nums := []int{2,7,11,15}
	
	fmt.Println(twoSum(nums,9))

}


