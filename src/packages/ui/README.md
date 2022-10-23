# Hangman Classic - Termbox Edition

Implementation of the [Termbox TUI Library](https://github.com/nsf/termbox-go) for the Hangman Classic game.

# How does Termbox work ?

## Termbox initialization

As said not in the non-existing Termbox documentation but in very small comments above the concerned functions in the code, Termbox must be initialized before any other built-in function is called. 

As Termbox must be closed when it is not useful anymore, we'll use a [defer function](https://www.educative.io/answers/what-is-the-defer-keyword-in-golang) to close it before the program ends.

Last part of the initialization is the indication of the Input Mode Termbox should be using when executing. Here, as we use the default InputEsc mode, there's no real need to go over that point any more than that.

    func main() {
	    err := Termbox.Init() 
	    if err != nil {
		    panic(err)
		}
		defer Termbox.Close()
		
		termbox.SetInputMode(termbox.InputEsc)
	}

   

## Termbox Canvas

This idea behind Termbox is simple: it uses the terminal as a canvas of cells. Think of these cells as the same thing as pixels, except that a cell is composed of one ASCII chararacter (even "empty" cells are made with whitespace characters).

![Here, the terminal is divided in cells of equal size.](https://i.ibb.co/8jzfSNz/canvas.png)
The coordinates of the first "E", for example, are x=1 and y=0.

The built-in function to register a character in a cell of the terminal is:

    func SetCell(x, y int, ch rune, fg, bg Attribute) 
    
    PARAMETERS:
    → x and y: coordinates as mentionned above.
    → ch: ASCII char that will populate the cell.
    → fg and bg: int values determining the color of both foreground (text)
      and background. 

To store the default color values of the terminal (applies to both background and foreground colors):

    const coldef = termbox.ColorDefault

   ## Printing the canvas
To print the canvas, a function must be called at the end of our current main, defining the content of one or several cells. For example, a very simple function:

    var char = 'A'
    func DrawChar(x, y int, ch rune) {
	    termbox.SetCell(x, y, ch, coldef, coldef)
	    termbox.Flush() 
	}
**NB:** Note the use of `termbox.Flush()` at the end of the drawing function: this function must be called in order to synchronize the internal cell buffer with the terminal.

If we call `DrawChar(2, 2, char)` at the end of our main, it will effectively display an 'A' character at coordinates [2:2] of our terminal.

If we want to do a visual interface based on the current size of the terminal, we can use `termbox.Size()`, which will return two int values: the first being the horizontal number of cells and the second being the vertical number of cells in the current terminal.

## Registering user key inputs
Termbox supports the use of all the keys of the keyboard and their combination (and also the mouse inputs, which we will not cover here).

This can be done using an infinite `for` loop, which would break for example when "ESC" key is hit. A program that would update the value of `char` based on the keyboard input or leave the program if the input is "ESC" would look like this:

    mainloop:
	    for  {
		    switch ev := termbox.PollEvent(); ev.Type {
		    case termbox.EventKey:
			    switch ev.Key {
				case termbox.KeyEsc:
					break mainloop
				default:
					if ev.Ch !=  0  {
						char = ev.Ch
					}
				}
			case termbox.EventError:
				panic(ev.Err)
			}
			
		DrawCell(2,  2,  char)
		}
Here, the `termbox.PollEvent()` function waits for an Event (i.e a mouseclick, a key hit, etc) and returns it. Then, if the Event is a key hit (`termbox.EventKey`), and that key is the "ESC" key (`termbox.KeyEsc`), the program breaks the loop.

Otherwise, the program changes the value of `char` for the character of the keyboard that has just been hit.

**NB:** Note the repetition of the `DrawCell` function at the end of the loop: it's needed if you want to update the terminal with the modifications that just occured.

# Hangman Implementation of Termbox

For now, it has been decided that the implementation of Termbox for our Hangman game would be almost the same as given by the [bonus instructions here](https://github.com/Lyon-Ynov-Campus/YTrack/tree/master/subjects/hangman/hangman-classic/piste-noire). 

The final visual the game should have is [displayed here](https://github.com/RathGate/Hangman-Classic/blob/termbox/assets/visuals.md).

![Main menu of the game](https://i.ibb.co/BZBFxSM/visual.png)

For now, the game itself does not work, nor does it register any input aside from "ESC", "ArrowLeft" and "ArrowRight". This will be done soon.

## Hangman-Termbox: Constants

Here, the game is not displayed dynamically according to the terminal screen size. The maximum size of the game screen should be 25 x 75 cells.

    const MAX_HEIGHT = 25
    const MAX_WIDTH = 71

## Handmade Functions

Several functions have been created to fit the needs we had for the game.


### `Fill(x, y, w, h int, cell termbox.Cell)`

 This function can be used to cover big areas with a specific Cell type. For example:

![Fill function](https://i.ibb.co/jbxLL8s/fill.png)

    Fill(3, 1, 5, 3, termbox.Cell{Ch: 'x'}
    
    PARAMETERS: 
    → x=3, y=1: the coordinates of the first cell to cover, (in purple on
     the example, note that the color is only here to illustrate).
    → w=5, h=3: size of the area to cover, here 5x3 cells.
    → termbox.Cell(ch: 'x'): the character 'x' is used to cover the cells.
      Note that the Cell structure has other attributes that can be used in 
      that function, for example foreground and background colors.
 ___

 ### `func  PrintText(x, y int, fg, bg termbox.Attribute, msg string)`

This function can be used to display a message in the terminal. The coordinates x and y are the position of the first letter of the string, the other will be displayed rightwards.

**NB:** Note that `\n` cannot be used here to jump lines and the terminal won't split the sentence if it goes beyond the border of the terminal. The splitting has to be done manually. 

![PrintText Function](https://i.ibb.co/YTz0Swb/hello-world.png)

    PrintText(1, 1, coldef, coldef, "Hello World ! ♥")
    
    PARAMETERS:
    → x=1, y=1: Coordinates of the first character of the string, here in purple.
    → fb, bg=Coldef: Text and Background color of the cells.
    → msg: String to be displayed on the given coordinates.
    
  ___

### `func PrintBox(x, y, w, h int, title, content string)`

An example might be useful to explain from the start:
![PrintBox Function.](https://i.ibb.co/r0trPX3/printbox.png)

This visual is the result of this:

    PrintBox(1, 1, 38, 3, "Suggest a word or a letter", "HELLOWORLD")
    
    PARAMETERS:
    → x=1, y=1: Coordinates of the first cell of the box, here in purple.
    → w=38, h=3: Size of the INSIDE of the box. This does NOT include the borders
      of the box (think of box-sizing: content-box in CSS notation), meaning
      that this box has actually a size of 40x5 cells.
    → title: string written on the top border of the box, at [x+2;y]. 
      Note that if the title is bigger than the size of the box, it simply won't
	  be printed.
	→ content: string written in the box. This is automatically formatted to be
	  printed both horizontally and vertically centered, meaning that if you want
	  to apply another format, it might be better to leave this parameter empty 
	  and fill the box in another way.
	  Note that for now, the string won't be splitted if it's bigger than the box.
