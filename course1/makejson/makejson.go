package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var name string
	var address string
	var person map[string]string
	person = make(map[string]string)

	fmt.Printf("Input your name: ")
	fmt.Scan(&name)
	fmt.Printf(name)
	person["name"] = name

	fmt.Printf("Input your address: ")
	fmt.Scan(&address)
	person["address"] = address

	personJson, _ := json.MarshalIndent(person,  "", "    ")

	fmt.Println(string(personJson))
}