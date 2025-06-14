package main

import (
	"fmt"
	// "strings"
)

type Animal struct {
	food string
	locomotion string
	sound string
}

func (animal Animal) Speak() {
	fmt.Println(animal.sound)
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func main() {
	var userAnimal string
	var userFuction string

	fmt.Printf("Input the animal name followed by function name, separeted by white space ' '. \n[Animal Options]: 'cow', 'bird', 'snake'. \n[Functions Options]: 'speak', 'eat', 'move'\n")

	cow := Animal{
		food: "grass",
		locomotion: "walk",
		sound: "moo",
	}
	bird := Animal{
		food: "worms",
		locomotion: "fly",
		sound: "peep",
	}
	snake := Animal{
		food: "mice",
		locomotion: "slither",
		sound: "hsss",
	}

	animals := map[string]Animal {
		"cow": cow,
		"bird": bird,
		"snake": snake,
	}

	functions := map[string]func(Animal) {
		"speak": Animal.Speak,
		"eat": Animal.Eat,
		"move": Animal.Move,
	}

	for true {
		fmt.Printf("> ")
		fmt.Scan(&userAnimal)
		fmt.Scan(&userFuction)
	
		

		animal, ok := animals[userAnimal]
		if !ok {
			fmt.Println("Invalid ANIMAL value, please enter a valid input.")
			continue
		}

		animalFunction, ok := functions[userFuction]
		if !ok {
			fmt.Println("Invalid FUNCTION value, please enter a valid input.")
			continue
		}

		fmt.Printf("%s %s: ", userAnimal, userFuction)
		animalFunction(animal)
	}
}

