package main

import (
	"hangman/packages/hangman"
	"hangman/packages/utils"
	"math/rand"
	"os"
	"time"
)

var hangmanData hangman.HangManData

func main() {
	// Initializes random seed and empties console before game starts.
	rand.Seed(time.Now().UnixNano())
	utils.ConsoleClear()

	// Checks for dict file in program args.
	if len(os.Args) == 1 {
		utils.PrintError("No dictionary file specified.")
	}
	filename := os.Args[1]

	// Launches game itself.
	hangmanData.InitGame(filename)
	for hangmanData.Attempts > 0 && !hangmanData.IsDiscovered() {
		hangman.NewRound(&hangmanData)
	}

	// TODO: here the game stops once attempts == 0 or
	// word had been discovered, without any message printed or anything.
}
