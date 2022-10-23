package hangman

import (
	"fmt"
	"hangman/packages/ascii"
	"hangman/packages/utils"
	"sort"
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
		fmt.Println(strings.Join(strings.Split(data.Word, ""), " "))
	}
	fmt.Print("\n")
}

// Launches a round of the game.
func NewRound(data *HangManData, charset *ascii.Charsets) {
	// Asks user for input and processes the answer.
	data.PrintStockedLetters()
	answer := utils.GetUserInput()
	processed := data.ProcessAnswer(data.FinalWord, answer)
	if processed == 0 {
		data.PrintWord(charset.Characters)
		// data.PrintStockedLetters()
	} else if processed < 0 {
		fmt.Printf("Not present in the word, %d attempts remaining.\n", data.Attempts)
		// data.PrintStockedLetters()
		ascii.PrintJose(charset.Jose, data.Attempts)
		fmt.Println()
	}
}

func (data *HangManData) RevealLetter(answer string) {

	for i, letter := range data.FinalWord {
		if answer == string(letter) {
			data.Word = data.Word[:i] + string(data.FinalWord[i]) + data.Word[i+1:]
		}
	}
}

func (data *HangManData) AddUsedLetters(letter string) bool {
	for _, char := range data.UsedLetters {
		if char == letter {
			fmt.Print("You already tried that letter.\n\n")
			return false
		}
	}
	data.UsedLetters = append(data.UsedLetters, letter)
	sort.Strings(data.UsedLetters)
	return true
}

func (data *HangManData) PrintStockedLetters() {
	fmt.Printf("Used letters: %v\n\n", strings.Join(data.UsedLetters, " "))
}

// Returns an int based on the points the player should lose or not.
// If 0, a letter had been discovered, if > 0, points should be lost.
// If 1, the right word has been suggested, the player has won.
func (data *HangManData) ProcessAnswer(word, answer string) int {
	// Answer = 1 character
	if len(answer) == 1 {
		if !data.AddUsedLetters(answer) {
			return 2
		}
		if strings.Contains(word, string(answer[0])) {
			data.RevealLetter(answer)
			return 0
		} else {
			data.LosePoints(1)
			return -1
		}
		// Answer = at least 2 characters.

	} else {
		if answer == "STOP" {
			SavePrompt(*data)
			return 0
		}
		if answer == word {
			data.Word = data.FinalWord
			return 1
		} else {
			data.LosePoints(2)
			return -2
		}
	}
}

func (data *HangManData) LosePoints(points int) {
	data.Attempts -= points
	if data.Attempts < 0 {
		data.Attempts = 0
	}
}
