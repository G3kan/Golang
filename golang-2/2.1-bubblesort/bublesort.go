package main

import (
	"fmt"
)

// Swap function: swaps elements at index i and i+1
func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

// BubbleSort function: sorts the array in ascending order
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	var nums []int
	var input int
	fmt.Println("Enter up to 10 integers (press Enter to stop early):")

	for len(nums) < 10 {
		fmt.Printf("Number #%d: ", len(nums)+1)
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input. Exiting.")
			break
		}
		nums = append(nums, input)
	}

	BubbleSort(nums)

	fmt.Print("Sorted array: ")
	for _, val := range nums {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
