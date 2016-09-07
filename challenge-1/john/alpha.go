package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

/*
John Pettigrew

Command to list the letters that are not being used in a file.
Inputs are not being checked for the sake of speed.
*/

const fileReadLength int = 10e+8

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	letters, err := markPresentLetters(file)
	if err != nil {
		panic(err)
	}

	for letter, letterVal := range *letters {
		if !letterVal {
			fmt.Printf("%c", letter)
		}
	}

}

func markPresentLetters(file *os.File) (*map[rune]bool, error) {

	letters := map[rune]bool{
		'a': false,
		'b': false,
		'c': false,
		'd': false,
		'e': false,
		'f': false,
		'g': false,
		'h': false,
		'i': false,
		'j': false,
		'k': false,
		'l': false,
		'm': false,
		'n': false,
		'o': false,
		'p': false,
		'q': false,
		'r': false,
		's': false,
		't': false,
		'u': false,
		'v': false,
		'w': false,
		'x': false,
		'y': false,
		'z': false,
	}

	var b []byte
	var str string

	for {
		b = make([]byte, fileReadLength)
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return &map[rune]bool{}, err
		}

		str = string(b[:fileReadLength])

		for _, runeVal := range str {

			letters[runeVal] = true
		}
	}

	return &letters, nil

}
