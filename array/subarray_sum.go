package main

import "fmt"

/*
Given an integer array, check if it contains a subarray having zero-sum.
Input:  { 3, 4, -7, 3, 1, 3, 1, -4, -2, -2 }
The subarrays with a sum of 0 are:
{ 3, 4, -7 }
{ 4, -7, 3 }
{ -7, 3, 1, 3 }
{ 3, 1, -4 }
{ 3, 1, 3, 1, -4, -2, -2 }
{ 3, 4, -7, 3, 1, 3, 1, -4, -2, -2 }
*/

// O(n²) Brute Force
func zeroSumList(arr []int) {
	found := false

	for i := 0; i < len(arr); i++ {
		sum := 0

		for j := i; j < len(arr); j++ {
			sum += arr[j]

			if sum == 0 {
				fmt.Printf("sublish exists :%v\n", arr[i:j+1])

				found = true
			}
		}
	}

	if !found {
		fmt.Println("sublish does not exists")
	}
}

func main() {

	nums := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}

	//zeroSumList(nums)

	hashZeroSumSubarray(nums)

}

// O(n) HashMap
func hashZeroSumSubarray(nums []int) {

	seen := make(map[int]int)

	sum := 0

	seen[0] = -1

	for i, num := range nums {

		sum += num

		if prevIndex, found := seen[sum]; found {

			fmt.Println(
				"Zero sum subarray:",
				nums[prevIndex+1:i+1],
			)
		}

		seen[sum] = i
	}
}
