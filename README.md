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
Usage:
```
Usage:
  go-passphrase [flags]

Flags:
  -c, --capitalize         Capitalize the first letter of the words
  -f, --half               Cut in half the words
  -h, --help               help for go-passphrase
  -s, --separator string   Character separating the words (default "-")
  -w, --words int          Number of words in the passphrase (default 3)
```

## How to install
Download the latest executable for your system in the [releases](https://github.com/Tashima42/go-passphrase/releases) page