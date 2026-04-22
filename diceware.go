/*
Get-passphrase generates a word sequence from the EFF Diceware Large Word List

By default, it generates and displays a 6-word sequence, but this can be modified
with some optional command line parameters.

Usage:

  diceware [-num=<# number of phrases to generate] [-len=<# of words per phrase>] [-wordfile=<path to word list file>] [-sep=<char>] [-suffix=<# random chars>] [-entropy]

Flags:

  -num = Number of phrases to generate. Default is 1.

  -len = Length of passphrase (number of words).  Default is 6.
 
  -wordfile = Path to external word list file.  Format must be either one
    word per line, or space-separated pair of "index word" per line.  Total
    of 7776 lines. Default is to use the internal EFF long word list.
 
  -sep = Separator character to use between words.  Default is "-".

  -suffix = Number of random alphanumeric characters (no lookalikes) to append
    after the words, joined with -sep.  Default is 0 (disabled).
    Useful for hybrid passphrases: e.g. -len=4 -suffix=4 yields roughly the
    same entropy as a pure 6-word phrase, but is shorter overall.

  -entropy = Print the estimated entropy (in bits) alongside each passphrase.
    Entropy is estimated assuming the default EFF word list; results may differ
    for custom word lists.

*/
package main

import (
  "local/diceware/eff"
  "fmt"
  "flag"
  "math"
  "strings"
  "os"
)

func main(){
  var phraseLength,numberOfEntries,suffixLength int
  var wordFile,wordSeparator string
  var passPhrase strings.Builder
  var debug,showEntropy bool

  flag.IntVar(&phraseLength,"len",6,"Number of words per phrase")
  flag.IntVar(&numberOfEntries,"num",1,"Number of phrases to generate")
  flag.IntVar(&suffixLength,"suffix",0,"Number of random alphanumeric chars to append (hybrid mode)")
  flag.StringVar(&wordFile,"wordfile","","Path to wordlist file (1 word/line, 7776 total lines)")
  flag.StringVar(&wordSeparator,"sep","-","Word separator character/string")
  flag.BoolVar(&showEntropy,"entropy",false,"Print estimated entropy (bits) alongside each passphrase")
  flag.BoolVar(&debug,"d",false,"Print debug information")

  flag.Parse()

  if phraseLength < 1 {
    fmt.Fprintln(os.Stderr, "Error: -len must be at least 1")
    os.Exit(1)
  }
  if numberOfEntries < 1 {
    fmt.Fprintln(os.Stderr, "Error: -num must be at least 1")
    os.Exit(1)
  }
  if suffixLength < 0 {
    fmt.Fprintln(os.Stderr, "Error: -suffix must be 0 or greater")
    os.Exit(1)
  }

  if debug{
    fmt.Println("Phrase Length: ", phraseLength)
    fmt.Println("Wordfile: ", wordFile)
    fmt.Println("Word Separator: ", wordSeparator)
    fmt.Println("Suffix Length: ", suffixLength)
  }

  if wordFile != "" {
    err := eff.LoadWordFile(wordFile)
    if err != nil {
      fmt.Fprintln(os.Stderr,"Error loading word file!",err)
      os.Exit(1)
    }
  }

  // Entropy estimate: each diceware word from 7776-word list contributes log2(7776) bits;
  // each suffix char from the 55-char no-lookalike set contributes log2(55) bits.
  // Separators add no entropy.
  entropyBits := float64(phraseLength)*math.Log2(7776) + float64(suffixLength)*math.Log2(float64(len(eff.DefaultSuffixChars)))

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

    if suffixLength > 0 {
      suffix, err := eff.GetRandomChars(suffixLength, eff.DefaultSuffixChars)
      if err != nil {
        fmt.Fprintln(os.Stderr,"Failed to generate random suffix.",err)
        os.Exit(1)
      }
      passPhrase.WriteString(wordSeparator)
      passPhrase.WriteString(suffix)
    }

    if showEntropy {
      fmt.Printf("%-50s (~%.1f bits)\n", passPhrase.String(), entropyBits)
    } else {
      fmt.Println(passPhrase.String())
    }
    passPhrase.Reset()
  }

}


