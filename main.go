package main

import (
	//"hangman/packages/ascii"
	"fmt"
	"hangman/packages/hangman"
	"os"
)

func main() {
	filename := os.Args[1]
	//charset := ascii.GetCharset("standard.txt", 8)
	lines := hangman.ReadFile(filename)
	randomWord := hangman.RandomWord(lines)
	fmt.Println(randomWord)

	var hangmanData hangman.HangManData
	hangmanData.InitHangMan(randomWord)
	fmt.Println(hangmanData.Word)

	//ascii.PrintAsciiWord(randomWord, charset)
}
