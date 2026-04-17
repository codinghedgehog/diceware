# Diceware
## Description
This is a simple go program that generates Diceware passphrases for use as passwords.
## Usage

By default it'll generate a 6 word, dash-separated passphrase drawn from the EFF Large Wordlist (https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt).

Additional options are available:
```
  -d    Print debug information
  -len int
        Number of words per phrase (default 6)
  -num int
        Number of phrases to generate (default 1)
  -sep string
        Word separator character/string (default "-")
  -wordfile string
        Path to wordlist file (either 1 word per line or "index word" column per line, 7776 total lines)
```
