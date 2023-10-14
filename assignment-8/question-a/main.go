package main

import (
	"fmt"
	"os"
)

func removeFile(fname string) {

	err := os.Remove(fname)

	if err != nil {

		fmt.Println("file not found to remove")
		return
	}

	fmt.Println("file removed")
}

func main() {

	path := `C:\Users\ORR Training 7\Desktop\Day-2&3\assignment-8\question-a\a.txt`

	removeFile(path)

}
