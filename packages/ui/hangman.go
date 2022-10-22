// All the functions used specifically to display
// the menus of the hangman game.

package ui

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// Constants defining the size of the game screen.
// Here, 70x25 cells (see ../assets/visuals).
const MAX_HEIGHT = 25
const MAX_WIDTH = 70

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
	PrintBox(20, 4, 18, 3, "Attempts", fmt.Sprintf("%d", currentMenu))
	//  Words or letters
	PrintBox(40, 4, 28, 18, "Hangman", "")
	//  Words or letters
	PrintBox(0, 9, 38, 3, "Suggest a word or a letter", "HELLOWORLD")
	//  Word
	PrintBox(0, 14, 38, 3, "Word", "_ _ L L O _ O _ R _ L _")
	// Used letters
	PrintBox(0, 19, 38, 3, "Wrong letters", "B F P")
}

// Prints both "Welcome" and "Help" menu as they have the same structure.
// Here, content should be `WELCOME_LINES` or `HELP_LINES` string array.
func TextMenu(x, y, w, h int, title string, content []string) {
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
