// All the functions used specifically to display
// the menus of the hangman game.

package ui

import (
	"fmt"
	"hangman/packages/ascii"
	"hangman/packages/hangman"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"
)

// Constants defining the size of the game screen.
// Here, 71x25 cells (see ../assets/visuals).
const MAX_HEIGHT = 25
const MAX_WIDTH = 71

var data *hangman.HangManData
var charset *ascii.Charsets
var editbox EditBox
var FINAL_WORD string

// Defines lines in the Welcome and Help menus.
var WELCOME_LINES = []string{
	`Welcome to the Termbox version of Hangman !`,
	`"Enter" to confirm your chosen letter or word.`,
	``,
	`Enjoy the game ! ♥`,
}
var HELP_LINES = []string{
	`1. "ESC" to quit.`,
	`2. "Enter" to confirm your choice.`,
	`3. "Backspace" or "Del" to delete the last letter.`,
}

// Defines all three menus available and sets the
// index of the current menu to "Welcome"
var menus = []string{"Welcome", "Game", "Help"}
var currentMenu = 0

// Function reacting to arrowLeft and arrowRight key presses.
// If arrowLeft: currentMenu-- ; if arrowRight: currentMenu++.
func switchMenu(dir int) {
	currentMenu += dir
	// Menu directly after the 3rd one is the 1st one and vice versa.
	if currentMenu < 0 {
		currentMenu = 2
	} else if currentMenu > 2 {
		currentMenu = 0
	}
}

// Main function to display the hangman screen.
// Draws the upper part of the screen common to all three menus,
// then based on the value of currentMenu, displays the appropriate
// part of the game.
func redrawHangman(x, y, w, h int) {
	UpperMenu(x, y)
	switch currentMenu {
	case 0:
		TextMenu(0, 4, 68, 8, "Welcome", WELCOME_LINES)
	case 1:
		GameMenu()
	case 2:
		TextMenu(0, 4, 68, 8, "Help", HELP_LINES)
	}
	// Must always be done at the end of the main display function.
	termbox.Flush()
}

// Displays the upper part of the game ("ESC to quit etc. + "Welcome | Game etc.".)
func UpperMenu(x, y int) {
	message := `"ESC" to quit, "Left" or "Right" to switch tabs`
	// Prints the upper text and the box for the main menu.
	PrintText(x, y, coldef, coldef, message)
	PrintBox(x, y+1, 48, 1, "", "")

	// Position of the first letter of the menu in the canvas,
	// statically written for now. Used to center the menu horizontally
	// in its box.
	centeredX := 14
	var tempLen int        // Position of the start of the current word.
	var sep string = " | " // Separation between each word of the menu.

	for i, menu := range menus {
		// Print the menu name in red if its index == currentMenu.
		if i == currentMenu {
			PrintText(centeredX+tempLen, 2, colred, coldef, menu)
		} else {
			PrintText(centeredX+tempLen, 2, coldef, coldef, menu)
		}

		// Prints separator in between menu names if currentMenu != the last of the list.
		if i != 2 {
			PrintText(centeredX+tempLen+len(menu), 2, coldef, coldef, sep)
			tempLen += len(menu) + 3
		}
	}
}

// Prints all little boxes of the main part of the game.
// For now, all the values are static and not based on
// the real game.
func GameMenu() {
	// Welcome
	PrintBox(0, 4, 18, 3, "", "Welcome !")
	// Attempts
	PrintBox(20, 4, 18, 3, "Attempts", strconv.Itoa(data.Attempts))
	//  Words or letters
	PrintBox(40, 4, 28, 19, "Hangman", "")
	if data.Attempts != 10 {
		DrawJose(40, 4, 28, 19, charset.Jose, data.Attempts)
	}
	//  Words or letters
	editbox.PrintEditBox(0, 9, 38, 3)
	//  Word
	PrintBox(0, 14, 38, 3, "Word", strings.Join(strings.Split(data.Word, ""), " "))
	// Used letters
	PrintBox(0, 19, 38, 4, "Used letters", strings.Join(data.UsedLetters, " "))
}

// Prints both "Welcome" and "Help" menu as they have the same structure.
// Here, content should be `WELCOME_LINES` or `HELP_LINES` string array.
func TextMenu(x, y, w, h int, title string, content []string) {
	termbox.HideCursor()
	PrintBox(x, y, w, h, title, "")
	// Content:
	for i, line := range content {
		PrintText(x+1, y+2+i, coldef, coldef, fmt.Sprintf(" %v ", line))
	}

	// Clears the end of the game screen if the box did not take all
	// the remaining height of the game screen.
	lastLine := 2 + h + y
	for i := lastLine; i <= MAX_HEIGHT; i++ {
		Fill(0, i, MAX_WIDTH, 1, termbox.Cell{Ch: ' '})
	}
}

func LaunchTBGame(hangmanData *hangman.HangManData, charSets *ascii.Charsets) {
	data = hangmanData
	FINAL_WORD = data.FinalWord
	charset = charSets
	// Termbox Initialization, should be done before any other function call.
	// After successful initialization, the library must be finalized using 'Close' function,
	// which is done here with defer function.
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Sets termbox Input Mode, must be done but not really important for us here.
	termbox.SetInputMode(termbox.InputEsc)

	// Draws the initial canvas.
	redrawHangman(0, 0, MAX_WIDTH, MAX_HEIGHT)

	// mainloop:
	// 2 = player quitted; 1 = player lost; 0 = player won.
	switch RunGame(hangmanData) {
	case 1:
		RunResults(hangmanData, "loss")
	case 0:
		RunResults(hangmanData, "win")
	}
}

func RunGame(data *hangman.HangManData) (state int) {
	for data.Attempts > 0 {
		if ev := termbox.PollEvent(); ev.Type == termbox.EventKey {
			switch ev.Key {
			// Reaction to "escape" key.
			case termbox.KeyEsc:
				return 2
				// Reaction to "arrow left" key.
			case termbox.KeyArrowLeft:
				switchMenu(-1)
				// Reaction to "arrow right" key.
			case termbox.KeyArrowRight:
				switchMenu(1)
			}
			if currentMenu == 1 {
				switch ev.Key {
				case termbox.KeyBackspace, termbox.KeyBackspace2:
					editbox.DeleteRune()
				case termbox.KeyDelete, termbox.KeyCtrlD:
					editbox.DeleteRuneForward()
				case termbox.KeyEnter:
					data.ProcessAnswer(data.FinalWord, strings.ToUpper(string(editbox.text)), "termbox")
					// Empties the Editbox
					var temp EditBox
					editbox = temp
				default:
					if ev.Ch != 0 {
						editbox.InsertRune(ev.Ch)
					}

				}
			}
		} else if ev.Type == termbox.EventError {
			panic(ev.Err)
		}

		if data.IsDiscovered() {
			return 0
		}
		// Redraws the canvas after the user has interacted with the program.
		redrawHangman(0, 0, MAX_WIDTH, MAX_HEIGHT)
	}
	return 1
}

func GetCenteredPos(baseX, baseY, boxW, boxH, elemW, elemH int) (posX, posY int) {
	posX = baseX + (boxW / 2) - (elemW / 2)
	posY = baseY + (boxH / 2) - (elemH / 2)
	return posX, posY
}

func DrawJose(x, y, w, h int, joses [][]string, attempts int) {
	position := len(joses) - 1 - attempts
	posX, posY := GetCenteredPos(x, y, w, h, len(joses[0]), len(joses))
	for i, line := range joses[position] {
		PrintText(posX, posY+i, coldef, coldef, line)
	}
}
