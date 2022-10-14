package ascii

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// TODO: Error handling
func GetCharSet(filename string, charSize int) [][]string {
	// Opens charset file
	path := "assets/ascii/" + filename
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Parses into array
	re := regexp.MustCompile(fmt.Sprintf(`(?:\r\n)((.+\n){%d})`, charSize))
	arr := re.FindAllString(string(content), -1)

	return SplitCharSet(arr)
}

// Splits each char into an array of size [char.height]
// Returns a 2D array.
func SplitCharSet(charset []string) (result [][]string) {
	for _, char := range charset {
		splitted := strings.Split(char, "\r\n")
		// excludes arr[0] and arr[max] which are only composed
		// of \r\n characters.
		result = append(result, splitted[1:len(splitted)-1])
	}
	return result
}

// Translate a char into its ASCII art equivalent.
// TODO: Error handling
func GetAsciiChar(char rune, charset [][]string) []string {
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
func PrintAsciiWord(word string, charset [][]string) {
	asciiWord := MakeAsciiWord(word, charset)

	if asciiWord == nil {
		fmt.Println("Nothing to see here:\nWord is empty.")
		return
	}
	for _, line := range asciiWord {
		fmt.Println(line)
	}
}
