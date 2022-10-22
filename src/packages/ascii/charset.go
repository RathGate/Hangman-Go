package ascii

import (
	"fmt"
	"hangman/packages/utils"
	"os"
	"regexp"
	"strings"
)

type Charsets struct {
	Jose       [][]string
	Characters [][]string
}

/* BASE EXPLANATION :
→ A charset is an array of size 95 containing all ASCII printable characters (32-127).
Each character is itself an array of all the lines composing the character.
→ A charset is therefore a 2D array. */

// Returns the 2D array charset from the specified filename.
// charSize is the number of lines composing the character,
// here, all three given files share a charSize of 8.
func GetCharset(path string, charsetSize, charSize int) [][]string {
	// Opens charset file

	content, err := os.ReadFile(path)
	if err != nil {
		utils.PrintError(err.Error())
	}

	// Splits the content of the file into an array of 95 characters.
	re := regexp.MustCompile(fmt.Sprintf(`(?:\r\n)((.+\n){%d})`, charSize))
	tempCharset := re.FindAllString(string(content), -1)
	if len(tempCharset) != charsetSize {
		utils.PrintError(fmt.Sprintf("load %v: invalid file format. \nThe file should contain %v characters and the system detected %v character(s).", path, charsetSize, len(tempCharset)))
	}

	// Splits each character of the preceding array into an array of lines.
	charset, err := SplitCharset(tempCharset)
	if err != nil {
		utils.PrintError(fmt.Sprintf("load %v: invalid file format. \nIrregular line width at %v.", path, err.Error()))
	}
	return charset
}

// Splits each character of an array of characters into an array of lines.
// Returns a 2D array.
func SplitCharset(tempCharset []string) (charset [][]string, err error) {

	for i, char := range tempCharset {
		// Separator is "\r\n" = the combination of characters
		// used to go to the next line (carriage return + newline)
		temp := strings.Split(char, "\r\n")
		// Excludes first and last elements of the array (which are empty).
		splitted := temp[1 : len(temp)-1]

		if !CharIsValid(splitted) {
			return nil, fmt.Errorf("char '%v', between line %d and %d", string(byte(i+32)), i*9+2, (i+1)*9)
		}
		charset = append(charset, splitted)
	}
	return charset, nil
}

// Checks if all lines composing a character are equal in size.
func CharIsValid(char []string) bool {
	lineWidth := len(char[0])
	for _, line := range char {
		if len(line) != lineWidth {
			return false
		}
	}
	return true
}

// Translates a characters into its ASCII art equivalent.
func GetAsciiChar(char rune, charset [][]string) []string {
	// Ignores characters which are not in the ASCII table.
	if char < 32 || char > 127 {
		return nil
	}
	index := int(char) - 32
	return charset[index]
}

// Translates an entire word into its ASCII art equivalent.
func MakeAsciiWord(word string, charset [][]string) []string {
	var result = make([]string, len(charset[0]))
	if len(word) == 0 {
		return nil
	}

	for _, letter := range word {
		for j, line := range GetAsciiChar(letter, charset) {
			result[j] += line
		}
	}
	return result
}

// Prints the ASCII word line by line.
func PrintAscii(word string, charset [][]string) {
	asciiWord := MakeAsciiWord(word, charset)

	// Checks if word is nil or empty.
	if asciiWord == nil || len(asciiWord[0]) == 0 {
		fmt.Println("Nothing to see here:\nWord is empty.")
		return
	}
	for _, line := range asciiWord {
		fmt.Println(line)
	}
}

func PrintJose(joses [][]string, attempts int) {
	position := len(joses) - 1 - attempts
	for _, line := range joses[position] {
		fmt.Println(line)
	}
}
