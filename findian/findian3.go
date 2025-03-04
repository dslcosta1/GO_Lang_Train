package main

/*********************************
PROBLEM:
// Write a program which prompts the user to enter a string. The program searches
// through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program
// should print “Found!” if the entered string starts with the character ‘i’, ends
// with the character ‘n’, and contains the character ‘a’. The program should print
// “Not Found!” otherwise. The program should not be case-sensitive, so it does
// not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example
// entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
// The program should print “Not Found!” for the following
// strings, “ihhhhhn”, “ina”, “xian”.

// Submit your source code for the program,
// “findian.go”.
*********************************/

/*********************************
NOTES:
Source: https://www.linkedin.com/pulse/mastering-golang-series-part-1-how-read-data-from-file-wang-ph-d-
Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object,
creating another object (Reader or Writer) that also implements the interface
but provides buffering and some help for textual I/O. Package fmt implements
formatted I/O with functions analogous to C's printf and scanf.
The format 'verbs' are derived from C's but are simpler.

fmt.Scanln(&user_input)
Scanln scans text read from standard input, storing successive space-separated
values into successive arguments. It stops scanning at a newline and after
the final item there must be a newline or EOF. Scanln can only take one line
from standard input, and it is not suitable for multiple-line input.

bufio.NewScanner(os.Stdin)
Scan advances the Scanner to the next token, which will then be available
through the Bytes or Text method. It returns false when the scan stops,
either by reaching the end of the input or an error. After Scan returns false,
the Err method will return any error that occurred during scanning,
except that if it was io.EOF, Err will return nil.

In general, always try to use the bufio.NewScanner for collecting input
from the console.
*********************************/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Buffered IO scanner for terminal input
	scanner := bufio.NewScanner(os.Stdin)
	// Prompt
	fmt.Println("Enter Something Weird:")
	// Scans the line from Stdin(Console)
	scanner.Scan()
	// Holds the string that scanned
	user_input := scanner.Text()

	// trim new lines; replace all spaces with no spaces; characters to lower case
	user_input = strings.Trim(user_input, "\n")
	user_input = strings.ReplaceAll(user_input, " ", "")
	user_input = strings.ToLower(user_input)

	// Problem Domain Vars
	starts_with_i := strings.HasPrefix(user_input, "i")
	ends_with_a_n := strings.HasSuffix(user_input, "n")
	contains_an_a := strings.ContainsAny(user_input, "a")

	// Compare user_input against Problem Domain
	if starts_with_i && ends_with_a_n && contains_an_a {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}