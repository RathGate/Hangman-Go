package main

import (
	"hangman/packages/game"
	"hangman/packages/utils"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	utils.ConsoleClear()

	var gameData game.HangManData
	args := os.Args[1:]

	if len(args) == 0 {
		utils.PrintError("No dictionary file specified.")
	}

	lines := game.ReadFile(args[0])
	randomWord := game.RandomWord(lines)

	gameData.InitGame(randomWord)
	for gameData.Attempts > 0 && !gameData.IsDiscovered() {
		game.NewRound(&gameData)
	}

	// TODO: here the game stops once attempts == 0 or
	// word had been discovered, without any message printed or anything.
}
