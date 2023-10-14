package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrCannotReadFile = errors.New("file cannot be read or not found")

func main() {

	file := `C:\Users\ORR Training 7\Desktop\Day-2&3\assignment-13\sample.txt`

	bytes, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(ErrCannotReadFile)
		return
	}

	str := string(bytes)

	fmt.Println(str)

	slc := strings.Fields(str)

	fmt.Println(slc)

	fmt.Println("Count of words are", len(slc))

	findToRemove := os.Remove(file)

	// Deleting File
	if findToRemove != nil {
		fmt.Println("file not found to delete")
		return
	}
	fmt.Println("file deleted")

}
