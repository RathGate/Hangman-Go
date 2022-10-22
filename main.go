package main

import (
	"fmt"
	"hangman/packages/ascii"
	"hangman/packages/hangman"
	"hangman/packages/ui"
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
	charSets.Jose = ascii.GetCharset("assets/jose/"+*joseFile, 10, 7)
	hangmanData.InitGame(*dictFile, *uiMode, charSets.Characters)

	if *uiMode == "termbox" {
		ui.LaunchTBGame()
	} else {
		for hangmanData.Attempts > 0 && !hangmanData.IsDiscovered() {
			hangman.NewRound(&hangmanData, &charSets)
		}

		if *uiMode == "asciiArt" {
			charSets.Characters = ascii.GetCharset("assets/ascii/"+*charFile, 95, 8)
			fmt.Println("success")
		}
		if hangmanData.Attempts <= 0 {
			fmt.Println("Poor José lost his head! Be ashamed.")
		} else if hangmanData.IsDiscovered() {
			fmt.Println("Congrats! You win.")
		}
	}

	// TODO: here the game stops once attempts == 0 or
	// word had been discovered, without any message printed or anything.
}
