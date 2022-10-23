package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hangman/packages/utils"
	"os"

	"github.com/logrusorgru/aurora"
)

func SavePrompt(data HangManData) {
	fmt.Print(aurora.BgBrightCyan("\nAre you sure you wanna quit ?"))
	fmt.Print("\n\n")
	fmt.Println("1 // Save and Quit")
	fmt.Println("2 // Quit without saving")

	answer := GetSaveInput()
	fmt.Println()
	switch answer {
	case 0:
		return
	case 1:
		SaveGame(data)
		os.Exit(1)
	case 2:
		os.Exit(1)
	}
}

func SaveGame(data HangManData) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		utils.PrintError("Something went wrong when converting data into json.\nPlease try again.")
	}
	err = os.WriteFile("savefiles/save.json", file, 0644)
	if err != nil {
		utils.PrintError("Something went wrong when saving data into json file.\nPlease try again.")
	}
}
func GetSaveInput() int {
	var answer int
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print("\n0// Back\n")
	for {
		fmt.Printf("→ Choose: ")
		if _, err := fmt.Fscanln(stdin, &answer); err != nil || answer < 0 || answer > 2 {
			utils.DiscardBuffer(stdin)
			fmt.Print("Not a valid answer. [0-1-2]\n\n")
			continue
		}
		break
	}
	return answer
}

func (data *HangManData) LoadFromSave(filename string) {
	content, err := os.ReadFile("savefiles/" + filename)
	if err != nil {
		utils.PrintError(err.Error())
	}
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		utils.PrintError(err.Error())
	}
}
