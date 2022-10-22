package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Custom error function similar to log.Fatal(err), but easier to read.
// Prints the error message and stops the program.
func PrintError(message string) {
	fmt.Printf("%v ", aurora.Bold(aurora.BgBrightRed("→→ ERROR ←←")))
	fmt.Print(message + "\n\n")
	os.Exit(1)
}

// ANSI escape code to empty the console.
func ConsoleClear() {
	fmt.Print("\033[H\033[2J")
}

// Empties the byte buffer between each input asked.
func DiscardBuffer(r *bufio.Reader) {
	r.Discard(r.Buffered())
}

// Asks the user for an input and checks if correct.
func GetUserInput() string {
	var answer string
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("→ Choose: ")
		if _, err := fmt.Fscanln(stdin, &answer); err != nil || !IsAlpha(answer) {
			DiscardBuffer(stdin)
			fmt.Print("Suggestion must be letter-only.\n\n")
			continue
		}
		break
	}
	return strings.ToUpper(answer)
}

// Checks if string is ONLY composed of alphabetical characters.
func IsAlpha(str string) bool {
	return regexp.MustCompile(`^[A-Za-z]+$`).MatchString(str)
}

func InitFlags() (dictPtr, josePtr, savePtr, modePtr, charsetPtr *string) {
	dictPtr = flag.String("dict", "words.txt", "Name of the dictionary file in use.")
	josePtr = flag.String("jose", "hangman.txt", "Name of the hangman ASCII art file in use.")
	savePtr = flag.String("save", "none", "Loads or not the .json save file before launching the game.")
	modePtr = flag.String("mode", "vanilla", "UI style that should be used.\n['vanilla', 'asciiArt', 'termbox']")
	charsetPtr = flag.String("charset", "standard.txt", "Name of the ASCII charset file in use.")

	flag.Parse()

	return dictPtr, josePtr, savePtr, modePtr, charsetPtr
}
