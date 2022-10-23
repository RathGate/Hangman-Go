package ui

import (
	"fmt"
	"hangman/packages/hangman"

	"github.com/nsf/termbox-go"
)

var LOSS_LINES = []string{
	fmt.Sprintf("Too bad... The word was '%v'.\n", FINAL_WORD),
	"\n",
	"Poor José lost his head! Be ashamed.",
}
var WIN_LINES = []string{
	fmt.Sprintf("You had it ! The word was '%v'.\n", FINAL_WORD),
	"\n",
	"Poor José won't lose his head this time. Congrats ! ♥",
}

func DrawResults(x, y, w, h int, result string) {
	PrintBox(0, 0, w, 1, "", "R E S U L T S")
	if result == "loss" {
		TextMenu(x, y+3, w, 17, "", LOSS_LINES)
	} else {
		TextMenu(x, y+3, w, 17, "", WIN_LINES)
	}

	DrawJose(0, 11, w, 11, charset.Jose, data.Attempts)
	termbox.Flush()
}

func RunResults(data *hangman.HangManData, result string) {
	DrawResults(0, 0, MAX_WIDTH-2, MAX_HEIGHT-2, result)
	for {
		if ev := termbox.PollEvent(); ev.Type == termbox.EventKey {
			switch ev.Key {
			// Reaction to "escape" key.
			case termbox.KeyEsc:
				return
			}
		} else if ev.Type == termbox.EventError {
			panic(ev.Err)
		}

		// Redraws the canvas after the user has interacted with the program.
		DrawResults(0, 0, MAX_WIDTH-2, MAX_HEIGHT-2, result)
	}
}
