package passphrase

import (
	"fmt"

	"github.com/mazen160/go-random"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PassphraseGenerator struct {
	Words     int
	Captalize bool
	Separator string
	HalfWords bool
}

func (pg *PassphraseGenerator) Generate() string {
	var word string

	for i := 0; i < pg.Words; i++ {
		randInt, err := random.GetInt(WordsListLength)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate random integer: %s", err))
		}
		selectedWord := WordsList[randInt]

		if pg.HalfWords {
			selectedWord = pg.HalfWord(selectedWord)
		}
		word += selectedWord

		if i != pg.Words-1 {
			word += pg.Separator
		}
	}

	if pg.Captalize {
		word = cases.Title(language.AmericanEnglish).String(word)
	}
	return word
}

func (pg *PassphraseGenerator) HalfWord(word string) string {
	half := len(word) / 2
	return word[0:half]
}
