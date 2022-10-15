package utils

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func PrintError(message string) {
	fmt.Printf("%v ", aurora.Bold(aurora.BgBrightRed("→→ ERROR ←←")))
	fmt.Print(message + "\n\n")
	os.Exit(1)
}
