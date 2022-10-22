package hangman

import (
	"fmt"
	"hangman/packages/ascii"
	"hangman/packages/utils"
	"strings"
)

// Checks if the word had been entirely uncovered.
func (data *HangManData) IsDiscovered() bool {
	return data.Word == data.FinalWord
}

// Prints the word with spaces between letters for readability.
func (data *HangManData) PrintWord(charset [][]string) {
	if len(data.Word) == 0 {
		utils.PrintError("Something went wrong, please try again.")
	}
	if charset != nil {
		ascii.PrintAscii(data.Word, charset)
	} else {
		for i, char := range data.Word {
			fmt.Print(string(char))
			if i != len(data.Word)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// Launches a round of the game.
func NewRound(data *HangManData, charset *ascii.Charsets) {
	// Asks user for input and processes the answer.
	answer := utils.GetUserInput()
	processed := ProcessAnswer(data.FinalWord, answer)

	if processed == 1 {
		// Word has been discovered
		data.Word = data.FinalWord
	} else if processed == 0 {
		// TODO: Uncover letters
		data.PrintWord(charset.Characters)
	} else {
		// The player had made a mistake = remove points.
		data.Attempts += processed
		if data.Attempts < 0 {
			data.Attempts = 0
		}

		fmt.Printf("Not present in the word, %d attempts remaining.\n", data.Attempts)
		ascii.PrintJose(charset.Jose, data.Attempts)
		fmt.Println()

	}
}

// Returns an int based on the points the player should lose or not.
// If 0, a letter had been discovered, if > 0, points should be lost.
// If 1, the right word has been suggested, the player has won.
func ProcessAnswer(word, answer string) int {
	// Answer = 1 character
	if len(answer) == 1 {
		if strings.Contains(word, string(answer[0])) {
			return 0
		} else {
			return -1
		}
		// Answer = at least 2 characters.
	} else {
		if answer == word {
			return 1
		} else {
			return -2
		}
	}
}
