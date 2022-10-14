package main

import (
	"hangman/ascii"
)

func main() {
	// utils.PrintError("bonjour")
	charSet := ascii.GetCharset("shadow.txt", 8)
	ascii.PrintAsciiWord("Stay hydrated bois", charSet)

}
