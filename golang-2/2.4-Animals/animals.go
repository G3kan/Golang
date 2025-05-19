package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface with required methods
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct
type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

// Bird struct
type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

// Snake struct
type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.Fields(scanner.Text())
		if len(input) != 3 {
			fmt.Println("Invalid command format. Use 'newanimal <name> <type>' or 'query <name> <info>'")
			continue
		}

		command, name, param := input[0], input[1], input[2]

		switch command {
		case "newanimal":
			var a Animal
			switch param {
			case "cow":
				a = Cow{}
			case "bird":
				a = Bird{}
			case "snake":
				a = Snake{}
			default:
				fmt.Println("Unknown animal type. Use 'cow', 'bird', or 'snake'.")
				continue
			}
			animals[name] = a
			fmt.Println("Created it!")

		case "query":
			animal, found := animals[name]
			if !found {
				fmt.Println("Animal not found.")
				continue
			}
			switch param {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query. Use 'eat', 'move', or 'speak'.")
			}

		default:
			fmt.Println("Unknown command. Use 'newanimal' or 'query'.")
		}
	}
}
