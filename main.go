package main

import "fmt"

func main() {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //a slice of ints from 0-10

	for i := 0; i < len(nums); i++ { //for loop iterating through all elements
		//if even
		if nums[i]%2 == 0 {
			fmt.Printf("%d is even\n", nums[i])
		} else { //if an element is odd
			fmt.Printf("%d is odd\n", nums[i])
		}
	}
}
