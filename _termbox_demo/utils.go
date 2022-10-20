// All the functions used globally across all demo programs.

package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// Constant defining the base color of the text and background.
const coldef = termbox.ColorDefault
const colred = 50

// Fills the canvas with a specified cell.
// x and y are the coordinates of the first cell to fill,
// and from here, the function replaces an area of w x h cells.
// See README.md for an example.
func Fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

// Prints a string from the coordinates [x:y] rightwards.
// See README.md for an example.
func PrintText(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

// Prints a box of size w x h at coordinates [x:y].
// See README.md for an example and the use of title and
// content parameters (optional).
func PrintBox(x, y, w, h int, title, content string) {
	// Corners:
	termbox.SetCell(x, y, '┌', coldef, coldef)
	termbox.SetCell(x, y+h+1, '└', coldef, coldef)
	termbox.SetCell(x+w+1, y, '┐', coldef, coldef)
	termbox.SetCell(x+w+1, y+h+1, '┘', coldef, coldef)
	// Sides:
	if h > 0 {
		Fill(x, y+1, 1, h, termbox.Cell{Ch: '│'})
		Fill(x+w+1, y+1, 1, h, termbox.Cell{Ch: '│'})
	}
	if w > 0 {
		Fill(x+1, y, w, 1, termbox.Cell{Ch: '─'})
		Fill(x+1, y+h+1, w, 1, termbox.Cell{Ch: '─'})
	}
	// Empty the box:
	Fill(x+1, y+1, w, h, termbox.Cell{Ch: ' '})
	// Title:
	if title != "" && len(title)+2 < w {
		PrintText(x+2, y, coldef, coldef, fmt.Sprintf(" %v ", title))
	}
	if content != "" && len(content)+2 < w {
		posX := x + (w/2 - len(content)/2) - len(content)%2
		posY := y + (h/2 + h%2)
		PrintText(posX, posY, coldef, coldef, fmt.Sprintf(" %v ", content))
	}
	termbox.Flush()
}
