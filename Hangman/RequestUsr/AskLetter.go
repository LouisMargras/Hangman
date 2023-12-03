package RequestUsr

import (
	"Hangman_Project/Verify"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskLetter(ListLetterUsed []string, ListASCII []string, CharList []string) string {
	// Initialisation du scanner pour lire une ligne depuis la console
	reader := bufio.NewReader(os.Stdin)

	// Boucle pour demander à l'utilisateur une lettre jusqu'à ce qu'une entrée valide soit obtenue
	for {
		fmt.Print("Choose a letter: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		// Supprimer les espaces et les sauts de ligne du début et de la fin de l'entrée
		Letter := strings.TrimSpace(input)

		// Vérifier si la lettre est unique et une majuscule
		if len(Letter) == 1 && 'A' <= Letter[0] && Letter[0] <= 'Z' {
			// Vérifier si la lettre a déjà été utilisée
			if Verify.VerifUsedLetter(Letter, ListLetterUsed, ListASCII, CharList) {
				fmt.Println("You cannot use this letter again!")
			} else {
				return Letter
			}
		} else {
			fmt.Println("Please enter a single uppercase letter.")
		}
	}
}
