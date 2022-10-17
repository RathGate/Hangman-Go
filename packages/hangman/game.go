package hangman

import (
	"bufio"
	"fmt"
	"hangman/packages/utils"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HangManData struct {
	Word      string // Word composed of '_', ex: H_ll_
	FinalWord string // Final word chosen by the program at the beginning. It is the word to find
	Attempts  int    // Number of attempts left
	//HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func (data *HangManData) InitGame(word string) {
	data.FinalWord = strings.ToUpper(word)
	data.Word = strings.Repeat("_", len(word))
	n := len(word)/2 - 1
	for i := 0; i < n; {
		r := rand.Intn(len(word))
		if data.Word[r] != byte('_') {
			continue
		} else {
			data.Word = data.Word[:r] + string(word[r]) + data.Word[r+1:]
			i++
		}
	}

	data.Attempts = 10

	fmt.Println("Good luck, you have 10 attempts.")
	data.PrintWord()
}

func ReadFile(filename string) []string {

	file, err := os.Open("assets/dict/" + filename)
	if err != nil {
		utils.PrintError(err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines
}

func RandomWord(lines []string) string {
	if len(lines) == 0 {
		utils.PrintError("Specified file is empty.")
	}

	rand.Seed(time.Now().UnixNano())
	return lines[rand.Intn(len(lines))]
}
