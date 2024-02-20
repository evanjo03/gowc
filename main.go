package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(os.Args)
	cFlag := flag.String("c", "", "Returns the number of bytes in a file.")
	lFlag := flag.String("l", "", "Returns the number of lines in a file.")
	wFlag := flag.String("w", "", "Returns the number of words in a file.")
	mFlag := flag.String("m", "", "Returns the number of chars in a file.")

	flag.Parse()

	if fileName := *cFlag; fileName != "" {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading the file: ", err)
			return
		}

		numBytes := len(data)
		fmt.Printf("%d %s\n", numBytes, fileName)
		return
	}

	if fileName := *lFlag; fileName != "" {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading the file: ", err)
			return
		}

		numLines := len(strings.Split(string(data), "\n"))
		fmt.Printf("%d %s\n", numLines, fileName)
		return
	}

	if fileName := *wFlag; fileName != "" {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading the file: ", err)
			return
		}

		numWords := len(strings.Fields(string(data)))
		fmt.Printf("%d %s\n", numWords, fileName)
		return
	}

	if fileName := *mFlag; fileName != "" {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading the file: ", err)
			return
		}
		numChars := utf8.RuneCountInString(string(data))
		fmt.Printf("%d %s\n", numChars, fileName)
		return
	}

	args, err := getFileFromArgs()
	if err != nil {
		fmt.Println(err)
	}

	fileName := args[0]
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}

	numBytes := len(data)
	numLines := len(strings.Split(string(data), "\n"))
	numWords := len(strings.Fields(string(data)))
	fmt.Printf("%d %d %d %s\n", numBytes, numLines, numWords, fileName)
}

func getFileFromArgs() ([]string, error) {
	args := os.Args[1:]
	if len(args) != 1 {
		return nil, errors.New("invalid arguments provided")
	}
	return args, nil
}
