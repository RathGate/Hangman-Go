package utils

import (
	"bufio"
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
