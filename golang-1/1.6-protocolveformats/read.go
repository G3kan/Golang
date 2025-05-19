package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	// Prompt the user for the file name

	fmt.Print("Enter the name of the text file: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a slice to store the structs
	var names []name

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 {
			// Create a struct and append it to the slice
			names = append(names, name{
				fname: strings.TrimSpace(parts[0]),
				lname: strings.TrimSpace(parts[1]),
			})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the names
	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.fname, n.lname)
	}
}
