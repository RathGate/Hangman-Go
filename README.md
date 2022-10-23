# Hangman Classic

Hangman game, realized by Eva Chibane and Marianne Corbel based on [these instructions](https://github.com/Lyon-Ynov-Campus/YTrack/tree/master/subjects/hangman/hangman-classic).

## About

This exercise must be submitted for notation before the 24th of October 2022. It is nevertheless possible to see some commits happening slightly after that date: no major change is planned, but some slight improvements in both code factorization and UI are not excluded to occur. 

It is up to you to take them into consideration, or not. ♪

If you're curious about our implementation of Termbox, a documentation page is available [here](https://github.com/RathGate/Corbel_Chibane_Hangman/tree/main/src/packages/ui), written by ourselves as well.

Our workflow has been documented [here](https://ynov890827.monday.com/boards/1131253654)

## How to use the program

This program has been tested and should work at least on Windows Terminal, both in and out of Visual Studio Code. Some tests should be done shortly in order to check it out soon.

This program uses flags to handle its different options:

    go run . -h

Consequence: 

    Usage of the program:
      -dict string
        Name of the dictionary file in use. (default "words.txt")
      -jose string
        Name of the hangman ASCII art file in use. (default "hangman.txt")
      -mode string
        UI style that should be used.
        ['vanilla', 'asciiArt', 'termbox'] (default "vanilla")
      -charset string
        Name of the ASCII charset file in use.
        Needs -mode=asciiArt to be considered. (default "standard.txt")
      -save string
        Loads or not the .json save file before launching the game. 
        (default "none")
You can use several of these flags, or none.

### Options available:

    dict: words.txt // words2.txt // words3.txt
    jose: hangman.txt
    mode: vanilla // asciiArt // termbox
    charset: standard.txt // thinkertoy.txt // shadow.txt
    save: save.json

**NB**:
1. For now, Termbox does handle save loading, but not saving into a file itself.
2. A saved game is not bound to its original game mode, ie. a game saved while is asciiArt mode can be loaded with -mode=termbox. But be aware of the point above.
3. Of course, to be taken into consideration, charset flag must come with -mode=asciiArt.

Enjoy the game ! ♥
