package main

import (
	"fmt"
	// "strings"
)

/*
Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print 
out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command 
or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of 
the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal 
and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the 
information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), 
which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method 
should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, 
Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when 
the user issues a query command.
*/

type Animal interface {
	Speak()
	Eat()
	Move()
}

//___________________________________________________________________________________________
type Cow struct {
	name string
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

//___________________________________________________________________________________________
type Bird struct {
	name string
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

//___________________________________________________________________________________________
type Snake struct {
	name string
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}


func main() {
	var userFuction string
	var userAnimal string
	var userArgument string
	

	var animalDB map[string]Animal

	animalDB = make(map[string]Animal)

	fmt.Printf("Input the animal name followed by function name, separeted by white space ' '. \n[Function Options]: 'newanimal', 'query'. \n[Animal Options]: 'cow', 'bird', 'snake'. \n[Functions Options]: 'speak', 'eat', 'move'\n")
	functions := map[string]func(Animal) {
		"speak": Animal.Speak,
		"eat": Animal.Eat,
		"move": Animal.Move,
	}

	for true {
		fmt.Printf("> ")
		fmt.Scan(&userFuction)
		fmt.Scan(&userAnimal)
		fmt.Scan(&userArgument)

		if userFuction == "newanimal" {
			switch userAnimal {
			case "cow":
				animalDB[userArgument] = Cow{userArgument}
			case "bird":
				animalDB[userArgument] = Bird{userArgument}
			case "snake":
				animalDB[userArgument] = Snake{userArgument}
			default:
				fmt.Println("Invalid ANIMAL value, please enter a valid input.")
				continue
			}

		} else {
			if userFuction == "query" {
				animal, ok := animalDB[userAnimal]
				if !ok {
					fmt.Println("Invalid ANIMAL_NAME value, please enter a valid input.")
					continue
				}

				animalFunction, ok := functions[userArgument]
				if !ok {
					fmt.Println("Invalid ARGUMENT value, please enter a valid input.")
					continue
				}
		
				animalFunction(animal)
			} else {
				fmt.Println("Invalid FUNCTION value, please enter a valid input.")
				continue
			}
		}
	}
}
