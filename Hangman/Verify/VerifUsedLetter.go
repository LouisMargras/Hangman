package Verify

import (
	"Hangman_Project/CreateList"
	"Hangman_Project/Print"
)

func VerifUsedLetter(Letter string, ListLetterUsed []string, ListASCII []string, CharList []string) bool { // Fonction qui vérifie si la lettre entrée a déjà été rentrée
	for i := range ListLetterUsed {
		if Letter == ListLetterUsed[i] {
			Message := "You cannot use this letter again!"
			ListWordASCII := CreateList.CreateASCIIWordList(Message, ListASCII, CharList)
			Print.PrintASCII(ListWordASCII)
			print("\n")
			return true
		}
	}
	return false
}
