package passphrase

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tashima42/go-passphrase/pkg/passphrase"
)

var (
	words         int
	captalize     bool
	separator     string
	halfWorlds    bool
	includeNumber bool
)

func InitCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-passphrase",
		Short: "Passphrase is a memorable passwords generator",
		Long: `Passphrases are passwords who are both easier to remember and more secure
                This tool is inspired by https://www.useapassphrase.com
                For more information, check the github page: https://github.com/Tashima42/go-passphrase`,
		Run: func(cmd *cobra.Command, args []string) {
			pg := passphrase.PassphraseGenerator{
				Words:         words,
				Captalize:     captalize,
				Separator:     separator,
				HalfWords:     halfWorlds,
				IncludeNumber: includeNumber,
			}
			fmt.Println(pg.Generate())
		},
	}

	rootCmd.Flags().IntVarP(&words, "words", "w", 3, "Number of words in the passphrase")
	rootCmd.Flags().BoolVarP(&captalize, "capitalize", "c", false, "Capitalize the first letter of the words")
	rootCmd.Flags().StringVarP(&separator, "separator", "s", "-", "Character separating the words")
	rootCmd.Flags().BoolVarP(&halfWorlds, "half", "f", false, "Cut in half the words")
	rootCmd.Flags().BoolVarP(&includeNumber, "number", "n", false, "Include a number at the end of one of the words")

	return rootCmd
}
