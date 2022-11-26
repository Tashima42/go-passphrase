package passphrase

import (
	"strings"
	"testing"
	"unicode"

	"github.com/tashima42/go-passphrase/cmd/passphrase"
)

func TestGeneratePassphrase(t *testing.T) {

	t.Run("one uncapitalized full word", func(t *testing.T) {
		pg := PassphraseGenerator{
			Words:     1,
			Captalize: false,
			Separator: "-",
			HalfWords: false,
		}
		generated := pg.Generate()
		if s := strings.Split(generated, pg.Separator); s[0] != generated {
			t.Errorf("Expected %s to be exactly equal to %s", s, &generated)
		}
		if f := generated[0:1]; !unicode.IsLower(f) {
			t.Errorf("Expected %s to be lower case", f)
		}
		if f := existsInArray(passphrase.WordsList, generated); f != true {
			if f := existsInArray(passphrase.WordsList, passphrase.HalfWord(generated); f == true) {
				t.Errorf("Expected %s to be a full word, not a half word", generated)
			}	else {
				t.Errorf("Expected %s to be in the words list", generated)
			}
		}
	})
	t.Run("three capitalized full words", func(t *testing.T) {
		pg := PassphraseGenerator{
			Words:     3,
			Captalize: true,
			Separator: "-",
			HalfWords: false,
		}
		generated := pg.Generate()
		if s := strings.Split(generated, pg.Separator); len(s) != pg.Words {
			t.Errorf("Expected %d length to be exactly equal to %d", s, pg.Words)
		}
		if f := generated[0:1]; unicode.IsLower(f) {
			t.Errorf("Expected %s to not be lower case", f)
		}
		if f := existsInArray(passphrase.WordsList, generated); f != true {
			if f := existsInArray(passphrase.WordsList, passphrase.HalfWord(generated); f == true) {
				t.Errorf("Expected %s to be a full word, not a half word", generated)
			}	else {
				t.Errorf("Expected %s to be in the words list", generated)
			}
		}
	})
}

func existsInArray(array [7776]string, element string) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}
