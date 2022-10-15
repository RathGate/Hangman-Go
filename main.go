package main

import (
	"hangman/packages/ascii"
	"hangman/packages/hangman"
	"os"
)

func main() {
	filename := os.Args[1]
	charset := ascii.GetCharset("standard.txt", 8)
	lines := hangman.ReadFile(filename)
	randomWord := hangman.RandomWord(lines)

	ascii.PrintAsciiWord(randomWord, charset)
}
