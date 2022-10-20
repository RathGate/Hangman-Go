package main

import "github.com/nsf/termbox-go"

func main() {
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

	// Sets the input the user can use to interact with the program.
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			// Reaction to "escape" key.
			case termbox.KeyEsc:
				break mainloop
				// Reaction to "arrow left" key.
			case termbox.KeyArrowLeft:
				switchMenu(-1)
				// Reaction to "arrow right" key.
			case termbox.KeyArrowRight:
				switchMenu(1)
			}
			// Error handle.
		case termbox.EventError:
			panic(ev.Err)
		}
		// Redraws the canvas after the user has interacted with the program.
		redrawHangman(0, 0, MAX_WIDTH, MAX_HEIGHT)
	}
}
