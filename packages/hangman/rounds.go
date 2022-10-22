package hangman

import (
	"fmt"
	"hangman/packages/utils"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Checks if the word had been entirely uncovered.
func (data *HangManData) IsDiscovered() bool {
	return data.Word == data.FinalWord
}

// Prints the word with spaces between letters for readability.
func (data *HangManData) PrintWord() {
	if len(data.Word) == 0 {
		utils.PrintError("Something went wrong, please try again.")
	}
	for i, char := range data.Word {
		fmt.Print(string(char))
		if i != len(data.Word)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n\n")

}

// Launches a round of the game.
func NewRound(data *HangManData) {
	// Asks user for input and processes the answer.
	answer := utils.GetUserInput()
	processed := data.ProcessAnswer(data.FinalWord, answer)

	if processed == 1 {
		// Word has been discovered
		data.Word = data.FinalWord
	} else if processed == 0 {
		data.RevealLetter(answer)
		data.PrintWord()
	} else {
		// The player had made a mistake = remove points.
		data.Attempts += processed
		if data.Attempts > 0 {
			fmt.Printf("Not present in the word, %d attempts remaining.\n", data.Attempts)
			fmt.Println(aurora.BgBrightCyan("Big José par ici."))
			fmt.Println()
		}
	}
}

func (data *HangManData) RevealLetter(answer string) {

	for i, letter := range data.FinalWord {
		if answer == string(letter) {
			data.Word = data.Word[:i] + string(data.FinalWord[i]) + data.Word[i+1:]
		}
	}
}

func (data *HangManData) AddUsedLetters(letter string) {
	for _, char := range data.UsedLetters {
		if char == letter {
			fmt.Println("You already tried that letter.")
			return
		}
	}
	data.UsedLetters = append(data.UsedLetters, letter)
}

func (data *HangManData) PrintStockedLetters() {
	fmt.Println("Used letters :", aurora.BgBlue(data.UsedLetters))
}

// Returns an int based on the points the player should lose or not.
// If 0, a letter had been discovered, if > 0, points should be lost.
// If 1, the right word has been suggested, the player has won.
func (data *HangManData) ProcessAnswer(word, answer string) int {
	// Answer = 1 character
	if len(answer) == 1 {
		data.AddUsedLetters(answer)
		data.PrintStockedLetters()
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
