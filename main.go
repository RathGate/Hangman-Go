package main

import (
	"fmt"
	"hangman/hangman"
	"os"
)

func main() {
	filename := os.Args[1]

	fmt.Println("Let's GOOO")

	lines := hangman.ReadFile(filename)
	fmt.Println(hangman.RandomWord(lines))
}
