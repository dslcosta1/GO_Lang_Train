package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type UserPrompt int
type Person struct {
	fname string
	lname string
}

const (
	MaxUserPromptLength = 45
	QuitCode            = "X"
	MapKeyName          = "name"
	MapKeyAddress       = "address"
	FieldSize           = 20
	BufferSize          = 42
	ListSize            = 10
)

const (
	PromptReadFilePath UserPrompt = iota
	PromptInvalidFilePath
	PromptSampleFilePath
	PromptExit
)

func GetUserPromptPadded(up UserPrompt, str string, strFmt string, newline bool) string {
	var paddedStr = str + strings.Repeat(" ", (MaxUserPromptLength-len(str))) + strFmt
	if newline {
		paddedStr += "\n"
	}
	return paddedStr
}

func GetUserPrompt(up UserPrompt) string {
	switch up {
	case PromptReadFilePath:
		return GetUserPromptPadded(up, "Please enter absolute file path:", "", false)
	case PromptInvalidFilePath:
		return GetUserPromptPadded(up, "Invalid path! Please enter a valid file path.", "", true)
	case PromptSampleFilePath:
		msg := "Eg: /dev/mooc/coursera/golang/week4/data.txt"
		return GetUserPromptPadded(up, msg, "", true)
	case PromptExit:
		return GetUserPromptPadded(up, "Enter 'X' to Exit", "", true)
	}
	return ""
}

func readValidInput(str string, regex *regexp.Regexp) (string, bool) {
	if str == QuitCode {
		os.Exit(0)
	}
	match := regex.Match([]byte(str))
	if !match {
		return "", false
	}
	return str, true
}

func readUntilValid(readPrompt UserPrompt, invalidPrompt UserPrompt, regex *regexp.Regexp) string {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf(GetUserPrompt(readPrompt))
	for scanner.Scan() {
		value, isValid := readValidInput(scanner.Text(), regex)
		if value != "" && isValid {
			return value
		}
		fmt.Printf(GetUserPrompt(invalidPrompt))
		if invalidPrompt == PromptInvalidFilePath {
			fmt.Printf(GetUserPrompt(PromptSampleFilePath))
		}
		fmt.Printf(GetUserPrompt(readPrompt))
	}
	return ""
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func getFileHandle(path string) (bool, *os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, nil, err
	}

	if fi, err := file.Stat(); err != nil || fi.IsDir() {
		return false, nil, err
	}
	return true, file, nil
}

func getFileAsString(file *os.File) (string, int, error) {
	fileinfo, err := file.Stat()
	if err != nil {
		return "", 0, err
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	byteCount, err := file.Read(buffer)
	if err != nil {
		return "", 0, err
	}

	return string(buffer), byteCount, nil
}

func getFileInChunkAsString(file *os.File) (string, int, error) {
	bufferSize := BufferSize
	buffer := make([]byte, bufferSize)
	str := ""
	for {
		byteCount, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return "", 0, err
		} else if err != nil && err == io.EOF {
			return str, byteCount, nil
		}
		content := string(buffer[:byteCount])
		str += content
	}
}

func getPersonSliceFromFile(file *os.File) []Person {
	reader := bufio.NewReader(file)
	var personSlice = make([]Person, 0, ListSize)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fields := strings.Split(string(line), " ")
		person := Person{fname: fields[0], lname: fields[1]}
		personSlice = append(personSlice, person)
	}
	return personSlice
}

func main() {
	fmt.Printf(GetUserPrompt(PromptExit))
	RegexForFilePath := regexp.MustCompile(`^(.+)/([^/]+)$`)

	filePath := readUntilValid(PromptReadFilePath, PromptInvalidFilePath, RegexForFilePath)
	exist, file, _ := getFileHandle(filePath)
	if !exist {
		fmt.Printf("%s: File DOES NOT exist.\n", filePath)
		os.Exit(0)
	}
	defer file.Close()

	persons := getPersonSliceFromFile(file)
	for _, person := range persons {
		fmt.Printf("fname: [%-20s]; lname: [%-20s]\n", person.fname, person.lname)
	}
}