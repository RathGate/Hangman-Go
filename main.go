package main

import (
	"hangman/ascii"
)

func main() {
	charSet := ascii.GetCharSet("thinkertoy.txt", 8)
	ascii.PrintAsciiWord("Hello World", charSet)
}
