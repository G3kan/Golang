package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type user struct {
	name   string
	adress string
}

func main() {

	// Create a map to hold the data
	data := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name:")
	inputname, _ := reader.ReadString('\n')
	inputname = strings.TrimSpace(inputname) // Remove any trailing newline or spaces

	fmt.Println("Enter address:")
	inputadress, _ := reader.ReadString('\n')
	inputadress = strings.TrimSpace(inputadress) // Remove any trailing newline or spaces

	// Create a user instance
	user1 := user{name: inputname, adress: inputadress}
	// Add the user to the map
	data[user1.name] = user1.adress
	// Print the map
	fmt.Println("Map data:", data)

	// Convert the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		os.Exit(1)
	}

	// Print JSON
	fmt.Println("JSON data:", string(jsonData))
}
