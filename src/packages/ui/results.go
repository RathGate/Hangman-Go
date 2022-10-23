package ui

import (
	"hangman/packages/hangman"

	"github.com/nsf/termbox-go"
)

func DrawResults(x, y, w, h int) {
	PrintBox(0, 0, w, h, "", "Welcome !")
}

func RunResults(data *hangman.HangManData) {
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
		DrawResults(0, 0, MAX_WIDTH-2, MAX_HEIGHT-2)
	}
}
