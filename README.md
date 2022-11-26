# go-passphrase
Passphrase generator.

## Passphrases are just better.
They are easier to remember and more secure

For more information, check https://www.useapassphrase.com

## How to use
Examples:   
1. Generate a passphrase with 4 captalized words:
    ```bash
    $ go-passphrase -w 4 -c
    > Patrol-Blandness-Professor-Five
    ```
1. Generate a passphrase with space as a separator:
    ```bash
    $ go-passphrase -s " "
    > geiger sadness founder
    ```
1. Generate a passphrase with half words:
    ```bash
    $ go-passphrase -f
    > geiger sadness founder
    ```
1. Generate a captalized passphrase with a number at the end of one of the words:
    ```bash
    $ go-passphrase -c -n
    > Create-Deluge1-Grandma
    ```
Usage:
```
Usage:
  go-passphrase [flags]

Flags:
  -c, --capitalize         Capitalize the first letter of the words
  -f, --half               Cut in half the words
  -n, --number             Include a number at the end of one of the words
  -h, --help               help for go-passphrase
  -s, --separator string   Character separating the words (default "-")
  -w, --words int          Number of words in the passphrase (default 3)
```

## How to install
Download the latest executable for your system and run.
* [MacOS Intel](https://github.com/Tashima42/go-passphrase/releases/download/v0.0.1-alpha/go-passphrase_darwin-amd64)
* [MacOS Apple Silicon](https://github.com/Tashima42/go-passphrase/releases/download/v0.0.1-alpha/go-passphrase_darwin-arm64)
* [Windows x86](https://github.com/Tashima42/go-passphrase/releases/download/v0.0.1-alpha/go-passphrase_windows-amd64)
* [Windows ARM](https://github.com/Tashima42/go-passphrase/releases/download/v0.0.1-alpha/go-passphrase_windows-arm64)
* [Linux x86](https://github.com/Tashima42/go-passphrase/releases/download/v0.0.1-alpha/go-passphrase_linux-amd64)
* [Linux ARM](https://github.com/Tashima42/go-passphrase/releases/download/v0.0.1-alpha/go-passphrase_linux-arm64)
