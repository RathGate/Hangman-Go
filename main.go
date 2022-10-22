package main

import (
	"fmt"
	"hangman/packages/ascii"
	"hangman/packages/hangman"
	"hangman/packages/utils"
	"math/rand"
	"time"
)

var hangmanData hangman.HangManData
var charSets ascii.Charsets

func main() {
	// Initializes random seed and empties console before game starts.
	dictFile, joseFile, _, uiMode, charFile := utils.InitFlags()
	rand.Seed(time.Now().UnixNano())
	// utils.ConsoleClear()
	// Launches game itself.
	charSets.Jose = ascii.GetCharset("assets/jose/"+*joseFile, 10, 7)
	if *uiMode == "asciiArt" {
		charSets.Characters = ascii.GetCharset("assets/ascii/"+*charFile, 95, 8)
		fmt.Println("success")
	}
	hangmanData.InitGame(*dictFile, charSets.Characters)
	for hangmanData.Attempts > 0 && !hangmanData.IsDiscovered() {
		hangman.NewRound(&hangmanData, &charSets)
	}

	// TODO: here the game stops once attempts == 0 or
	// word had been discovered, without any message printed or anything.
}
