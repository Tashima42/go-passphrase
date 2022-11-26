package passphrase

import (
	"strings"
	"testing"
	"unicode"
)

func TestGeneratePassphrase(t *testing.T) {
	t.Run("one word", func(t *testing.T) {
		pg := PassphraseGenerator{
			Words:     1,
			Captalize: false,
			Separator: "-",
			HalfWords: false,
		}
		generated := pg.Generate()
		if s := strings.Split(generated, pg.Separator); s[0] != generated {
			t.Errorf("Expected generated word %s to be exactly equal to %s", generated, s)
		}
	})
	t.Run("three words", func(t *testing.T) {
		pg := PassphraseGenerator{
			Words:     3,
			Captalize: true,
			Separator: "-",
			HalfWords: false,
		}
		generated := pg.Generate()
		if s := strings.Split(generated, pg.Separator); len(s) != pg.Words {
			t.Errorf("Expected number of words to be exactly equal to %d, got %d ", pg.Words, len(s))
		}
	})
	t.Run("three uncapitalized words", func(t *testing.T) {
		pg := PassphraseGenerator{
			Words:     3,
			Captalize: false,
			Separator: "-",
			HalfWords: false,
		}
		generated := pg.Generate()
		if isLower(strings.Join(strings.Split(generated, pg.Separator), "")) != true {
			t.Errorf("Expected %s to be lower case", generated)
		}
	})
	t.Run("two capitalized words", func(t *testing.T) {
		pg := PassphraseGenerator{
			Words:     2,
			Captalize: true,
			Separator: "-",
			HalfWords: false,
		}
		generated := pg.Generate()
		separatedWords := strings.Split(generated, pg.Separator)
		if len(separatedWords) != pg.Words {
			t.Errorf("Expected number of words to be exactly equal to %d, got %d ", pg.Words, len(separatedWords))
			return
		}
		if isLower(strings.Join(strings.Split(separatedWords[0], pg.Separator), "")) == true {
			t.Errorf("Expected %s to not be lower case", separatedWords[0])
		}
		if isLower(strings.Join(strings.Split(separatedWords[1], pg.Separator), "")) == true {
			t.Errorf("Expected %s to not be lower case", separatedWords[1])
		}
	})
	t.Run("one uncapitalized half word", func(t *testing.T) {
		pg := PassphraseGenerator{
			Words:     1,
			Captalize: false,
			Separator: "-",
			HalfWords: true,
		}
		generated := pg.Generate()
		if f := existsInArray(WordsList, pg.HalfWord(generated)); f == true {
			t.Errorf("Expected %s to be a full word, not a half word", generated)
		}
	})
}

func existsInArray(array [WordsListLength]string, element string) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
