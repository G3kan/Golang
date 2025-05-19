package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal struct with three attributes
type Animal struct {
	food       string
	locomotion string
	sound      string
}

// Methods for Animal
func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	// Predefined animals
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		// Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.ToLower(strings.TrimSpace(input))
		parts := strings.Split(input, " ")

		if len(parts) != 2 {
			fmt.Println("Please enter two words: [animal] [action]")
			continue
		}

		animalName := parts[0]
		action := parts[1]

		animal, found := animals[animalName]
		if !found {
			fmt.Println("Unknown animal:", animalName)
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Unknown action. Use 'eat', 'move', or 'speak'.")
		}
	}
}
