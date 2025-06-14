package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
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

func main() {
	var inputStr string
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		
		fmt.Println("Input animal name and method")
		
		fmt.Printf("> ")
		if scanner.Scan() {
			inputStr = scanner.Text()
		}
		splitStr := strings.Split(inputStr, " ")
		
		animalName := splitStr[0]
		method := splitStr[1]
		
		animal, err := factory(animalName)

		if err != nil {
			fmt.Println("animal does not exist, inform a existing animal (cow, snake or bird) ")
			continue
		}

		applyMethod(method, animal)

		fmt.Println("")
	}

}

func factory(animalName string) (Animal, error) {
	animalNameLower := strings.ToLower(animalName)
	switch animalNameLower {
	case "cow":
		return Animal{food: "grass", locomotion: "walk", noise: "moo"}, nil
	case "bird":
		return Animal{food: "worms", locomotion: "fly", noise: "peep"}, nil
	case "snake":
		return Animal{food: "mice", locomotion: "slither", noise: "hsss"}, nil
	default:
		err := errors.New("animal not found")
		return Animal{}, err
	}
}

func applyMethod(method string, animal Animal) {
	parseMethod := strings.ToUpper(method)
	switch parseMethod {
	case "EAT":
		animal.Eat()
	case "MOVE":
		animal.Move()
	case "SPEAK":
		animal.Speak()
	default:
		fmt.Println("METHOD DOES NOT EXISTS")
	}
}