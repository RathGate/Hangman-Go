package main

import (
	"fmt"
	"hangman/packages/ascii"
	"hangman/packages/hangman"
	"hangman/packages/ui"
	"hangman/packages/utils"
)

var hangmanData hangman.HangManData
var charSets ascii.Charsets

func main() {
	// Initializes random seed and empties console before game starts.
	dictFile, joseFile, saveFile, uiMode, charFile := utils.InitFlags()

	charSets.Jose = ascii.GetCharset("assets/jose/"+*joseFile, 10, 7)
	if *uiMode == "asciiArt" {
		charSets.Characters = ascii.GetCharset("assets/ascii/"+*charFile, 95, 8)
	}

	hangmanData.InitGame(*dictFile, *uiMode, charSets.Characters, *saveFile)

	if *uiMode == "termbox" {
		ui.LaunchTBGame(&hangmanData, &charSets)
	} else {
		for hangmanData.Attempts > 0 && !hangmanData.IsDiscovered() {
			hangman.NewRound(&hangmanData, &charSets)
		}
		if hangmanData.Attempts <= 0 {
			fmt.Println("Poor José lost his head! Be ashamed.")
		} else if hangmanData.IsDiscovered() {
			fmt.Println("Congrats! You win.")
		}
	}
}
