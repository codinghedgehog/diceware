/*
Get-passphrase generates a word sequence from the EFF Diceware Large Word List

By default, it generates and displays a 6-word sequence, but this can be modified
with some optional command line parameters.

Usage:

  diceware [-num=<# number of phrases to generate] [-len=<# of words per phrase>] [-wordfile=<path to word list file>] [-sep=<char>]

Flags:

  -num = Number of phrases to generate. Default is 1.

  -len = Length of passphrase (number of words).  Default is 6.
 
  -wordfile = Path to external word list file.  Format must be either one
    word per line, or space-separated pair of "index word" per line.  Total
    of 7776 lines. Default is to use the internal EFF long word list.
 
  -sep = Separator character to use between words.  Default is "-".

*/
package main

import (
  "local/diceware/eff"
  "fmt"
  "flag"
  "strings"
  "os"
)

func main(){
  var phraseLength,numberOfEntries int
  var wordFile,wordSeparator string
  var passPhrase strings.Builder
  var debug bool

  // Learning example for using flag directly, rather than passing variable references.
  //lengthPtr := flag.Int("length",6,"Number of words to generate")
  //wordFilePtr := flag.String("wordfile","","Path to wordlist file (1 word/line, 7776 total lines)")
  //wordSeparatorPtr := flag.String("sep","-","Word separator character/string")

  flag.IntVar(&phraseLength,"len",6,"Number of words per phrase")
  flag.IntVar(&numberOfEntries,"num",1,"Number of phrases to generate")
  flag.StringVar(&wordFile,"wordfile","","Path to wordlist file (1 word/line, 7776 total lines)")
  flag.StringVar(&wordSeparator,"sep","-","Word separator character/string")
  flag.BoolVar(&debug,"d",false,"Print debug information")

  flag.Parse()

  if debug{
    fmt.Println("Phrase Length: ", phraseLength)
    fmt.Println("Wordfile: ", wordFile)
    fmt.Println("Word Separator: ", wordSeparator)
  }

  if wordFile != "" {
    err := eff.LoadWordFile(wordFile)
    if err != nil {
      fmt.Fprintln(os.Stderr,"Error loading word file!",err)
      os.Exit(1)
    }
  }

  for _ = range numberOfEntries {
    wordList,err := eff.GetWords(phraseLength)
    if err != nil{
      fmt.Fprintln(os.Stderr,"Failed to generate passphrase.  Corrupt/improper word list?",err)
      os.Exit(1)
    }

    for i, w := range wordList {
      if i > 0 {
        passPhrase.WriteString(wordSeparator)
      }
      passPhrase.WriteString(w)
    }

    fmt.Println(passPhrase.String())
    passPhrase.Reset()
  }

}


