package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	fmt.Println("You will be prompted to enter animal and action separated by a single space, " +
		"you can stop the program by pressing Ctrl+C on your keyboard.")
	fmt.Println("Allowed animals: cow, bird, snake")
	fmt.Println("Allowed actions: eat, move, speak")

	fmt.Print("\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		animal, action, found := strings.Cut(input, " ")
		if !found {
			fmt.Println("Please, enter animal and action separated by a space. For example: cow eat")
			fmt.Print("> ")
			continue
		}

		var animalInstance Animal
		switch animal {
		case "cow":
			animalInstance = cow
		case "bird":
			animalInstance = bird
		case "snake":
			animalInstance = snake
		default:
			fmt.Println("Unknown animal. Allowed animals: cow, bird, snake")
		}

		switch action {
		case "eat":
			animalInstance.Eat()
		case "move":
			animalInstance.Move()
		case "speak":
			animalInstance.Speak()
		default:
			fmt.Println("Unknown action. Allowed actions: eat, move, speak")
		}

		fmt.Print("> ")
	}

	err := scanner.Err()
	if err != nil {
		fmt.Println("Error occurred while reading input: ", err)
		return
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}