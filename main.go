package main

import (
	"fmt"
	"hangman/packages/hangman"
	"hangman/packages/utils"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	utils.ConsoleClear()

	var hangmanData hangman.HangManData

	if len(os.Args) == 1 {
		utils.PrintError("No dictionary file specified.")
	}
	filename := os.Args[1]
	lines := hangman.ReadFile(filename)

	randomWord := hangman.RandomWord(lines)
	fmt.Println(randomWord)

	hangmanData.InitGame(randomWord)
	for hangmanData.Attempts > 0 && !hangmanData.IsDiscovered() {
		hangman.NewRound(&hangmanData)
	}

	// TODO: here the game stops once attempts == 0 or
	// word had been discovered, without any message printed or anything.
}
