package main

import (
	"fmt"
	"os"
	"strings"
)

type aPerson struct {
	fname string
	lname string
}

func main() {
	const stringSize int = 42

	var filename string
	var nb int = stringSize
	people := []aPerson{}
	fmt.Printf("Input the filename: ")
	fmt.Scan(&filename)

	file, _ := os.Open(filename)

	for nb == stringSize {
		barr := make([]byte, stringSize)
		nb, _ = file.Read(barr)

		personFullName := string(barr)
		personNames := strings.Split(personFullName, " ")

		

		person := aPerson{
			fname: personNames[0],
			lname: strings.Replace(personNames[1], "\n", "", 1),
		}

		people = append(people, person)
		fmt.Println(people)
	}

	for _, personName := range people {
		fmt.Printf("FirstName: %s | LastName: %s\n", personName.fname, personName.lname)
	}
}