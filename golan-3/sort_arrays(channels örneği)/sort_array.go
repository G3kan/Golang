package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Merge two sorted slices into one sorted slice
func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// Sort a subarray in a goroutine
func sortSubarray(wg *sync.WaitGroup, subarray []int, resultChan chan []int) {
	defer wg.Done()
	fmt.Println("Sorting subarray:", subarray)
	sort.Ints(subarray)
	resultChan <- subarray
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a series of integers separated by spaces:")

	scanner.Scan()
	input := scanner.Text()
	strNums := strings.Fields(input)

	if len(strNums) < 4 {
		fmt.Println("Please enter at least 4 integers.")
		return
	}

	// Convert input to a slice of integers
	nums := make([]int, 0, len(strNums))
	for _, s := range strNums {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid input:", s)
			return
		}
		nums = append(nums, num)
	}

	// Determine the size of each partition
	partSize := len(nums) / 4
	extra := len(nums) % 4

	var partitions [][]int
	start := 0
	for i := 0; i < 4; i++ {
		end := start + partSize
		if i < extra {
			end++
		}
		partitions = append(partitions, nums[start:end])
		start = end
	}

	var wg sync.WaitGroup
	resultChan := make(chan []int, 4)

	// Start goroutines to sort each partition
	for _, part := range partitions {
		wg.Add(1)
		go sortSubarray(&wg, part, resultChan)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(resultChan)

	// Collect sorted subarrays
	var sortedParts [][]int
	for sorted := range resultChan {
		sortedParts = append(sortedParts, sorted)
	}

	// Merge sorted subarrays
	merged := sortedParts[0]
	for i := 1; i < len(sortedParts); i++ {
		merged = merge(merged, sortedParts[i])
	}

	fmt.Println("Final sorted array:", merged)
}
