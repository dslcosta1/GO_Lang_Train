package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// MaxNameSize Max size of 20 is required, I'll use this constant to truncate strings
// instead of hardcoding 20 everywhere
const MaxNameSize = 20

type Name struct {
	fname string
	lname string
}

func main() {
	names := make([]Name, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the name of the text file: ")
	scanner.Scan()
	filename := scanner.Text()

	fmt.Println("Processing file ", filename, "...")
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error reading filename: ", err)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	fileScanner := bufio.NewScanner(file)
	line := 0
	for fileScanner.Scan() {
		line++
		firstName, lastName, found := bytes.Cut(fileScanner.Bytes(), []byte(" "))

		if !found {
			// fmt.Println("Space not found at line #", line, ".. skipping")
			// some people around the world have only one name
			// so instead of skipping, let's add them to the slice too
			names = append(names, Name{
				fname: truncate(string(fileScanner.Bytes())),
				lname: "",
			})
			continue
		}

		fn := truncate(string(firstName))
		ln := truncate(string(lastName))

		names = append(names, Name{
			fname: fn,
			lname: ln,
		})
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning file: ", err)
	}

	if file.Close() != nil {
		fmt.Println("Error closing file: ", err)
	}

	fmt.Println("\nResult: ")
	for idx, name := range names {
		fmt.Printf("Entry #%d: First name: %s, Last name: %s\n", idx, name.fname, name.lname)
	}
}

func truncate(s string) string {
	if len(s) > MaxNameSize {
		s = s[:MaxNameSize]
	}

	return s
}