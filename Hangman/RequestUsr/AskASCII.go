package RequestUsr

import (
	"bufio"
	"os"
)

func AskASCII() string {
	Scanner := bufio.NewScanner(os.Stdin)
	print("choose a model of ASCII (Standard : St, Shadow : Sh, Thinkertoy : Th):\n")
	Scanner.Scan()          // Lancement du Scanner
	ASCII := Scanner.Text() // Stockage du résultat du Scanner dans une variable
	if ASCII == "St" {
		return "standard.txt"
	} else if ASCII == "Sh" {
		return "shadow.txt"
	} else if ASCII == "Th" {
		return "thinkertoy.txt"
	} else {
		print("Your choice is not valid, try again!\n")
		return ""
	}
}
