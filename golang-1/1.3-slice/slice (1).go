package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){

	var numbers []int

	for {
		var input string
		fmt.Print("Give me number (to finish press 'X' ): ")
		fmt.Scanln(&input)

		
		if input == "X" {
			break
		}


		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not acceptable number")
			continue
		}

		
		numbers = append(numbers, num)

		
		sort.Ints(numbers)

		
		fmt.Println("Sorted:", numbers)
	}

	fmt.Println("Finished")
}
