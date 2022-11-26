package main

import (
	"fmt"
	"os"

	"github.com/tashima42/go-passphrase/cmd/passphrase"
)

func main() {
	rootCmd := passphrase.InitCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
