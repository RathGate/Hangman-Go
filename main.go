package main

import (
	"hangman/ascii"
)

func main() {
	charSet := ascii.GetCharSet("standard.txt", 8)
	charSet2 := ascii.GetCharSet("shadow.txt", 8)
	ascii.PrintAsciiWord("Ladies and Gentlemen", charSet)
	ascii.PrintAsciiWord("The Weeknd", charSet2)
}
