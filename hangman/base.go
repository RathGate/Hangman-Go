package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word   string // Word composed of '_', ex: H_ll_
	ToFind string // Final word chosen by the program at the beginning. It is the word to find
	//Attempts         int // Number of attempts left
	//HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func ReadFile(filename string) []string {

	fmt.Println(filename)
	file, err := os.Open("assets/dict/" + filename)
	check(err)

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
	//fmt.Println(rand.Intn(len(lines)))
	return lines[rand.Intn(len(lines))]
}