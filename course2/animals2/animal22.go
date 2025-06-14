package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

// Cow implements the Animal interface methods
func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird implements the Animal interface methods
func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake implements the Animal interface methods
func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		var command, name, animalType, queryType string

		// Print the prompt
		fmt.Print("> ")

		// Read the first part of the input (command, name, animalType)
		_, err := fmt.Scanf("%s %s %s\n", &command, &name, &animalType)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Handle the "newanimal" command
		if command == "newanimal" {
			var newAnimal Animal

			switch animalType {
			case "cow":
				newAnimal = Cow{name: name}
			case "bird":
				newAnimal = Bird{name: name}
			case "snake":
				newAnimal = Snake{name: name}
			default:
				fmt.Println("Unknown animal type:", animalType)
				continue
			}

			// Add the new animal to the map
			animals[name] = newAnimal

			// Confirm creation
			fmt.Println("Created it!")

		} else if command == "query" {
			// Clear the input buffer (if any leftover characters) by reading the next newline
			_, err := fmt.Scanln() // This will consume the newline left in the buffer
			if err != nil {
				fmt.Println("Error clearing buffer:", err)
			}

			// Read the query type (eat, move, speak)
			fmt.Print("> ")
			_, err = fmt.Scanln(&queryType)
			if err != nil {
				fmt.Println("Error reading query:", err)
				continue
			}

			// Find the animal by name
			animal, exists := animals[name]
			if !exists {
				fmt.Println("No such animal exists!")
				continue
			}

			// Handle the query based on the query type
			switch queryType {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query")
			}
		} else {
			// Handle invalid commands
			fmt.Println("Invalid command")
		}
	}
}