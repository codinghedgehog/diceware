package eff

import (
  "fmt"
)

type DicewareNumber struct {
  Digit1 int
  Digit2 int
  Digit3 int
  Digit4 int
  Digit5 int
}

func (dn DicewareNumber) valid() bool{
  return dn.Digit1 > 0 && 
        dn.Digit1 <= 6 &&
        dn.Digit2 > 0 &&
        dn.Digit2 <= 6 &&
        dn.Digit3 > 0 &&
        dn.Digit3 <= 6 &&
        dn.Digit4 > 0 &&
        dn.Digit4 <= 6 &&
        dn.Digit5 > 0 &&
        dn.Digit5 <= 6 


}

func NewDicewareNumber() *DicewareNumber {
  dn := new(DicewareNumber)
  dn.Reset()
  return dn
}


func (dn *DicewareNumber) Reset() {
  dn.Digit1 = 1
  dn.Digit2 = 1
  dn.Digit3 = 1
  dn.Digit4 = 1
  dn.Digit5 = 1
}

func (dn DicewareNumber) StringValue() string {

  if ! dn.valid() {
    panic("Invalid DicewareNumber representation! Need to init first?")
  }

  return fmt.Sprintf("%d%d%d%d%d",dn.Digit1,dn.Digit2,dn.Digit3,dn.Digit4,dn.Digit5)

}

func (dn DicewareNumber) IntValue() int {
  if ! dn.valid() {
    panic("Invalid DicewareNumber representation! Need to init first?")
  }
  numericValue := dn.Digit1 * 10000 + dn.Digit2 * 1000 + dn.Digit3 * 100 + dn.Digit4 * 10 + dn.Digit5
  return numericValue

}

func (dn *DicewareNumber) Inc() {
  if ! dn.valid() {
    panic("Invalid DicewareNumber representation! Need to init first?")
  }
  carryOver := false
  if dn.Digit5 == 6 {
    dn.Digit5 = 1
    carryOver = true
  } else {
    dn.Digit5++
  }

  if carryOver {
    carryOver = false
    if dn.Digit4 == 6 {
      dn.Digit4 = 1
      carryOver = true
    } else {
      dn.Digit4++
    }
  }

  if carryOver {
    carryOver = false
    if dn.Digit3 == 6 {
      dn.Digit3 = 1
      carryOver = true
    } else {
      dn.Digit3++
    }
  }

  if carryOver {
    carryOver = false
    if dn.Digit2 == 6 {
      dn.Digit2 = 1
      carryOver = true
    } else {
      dn.Digit2++
    }
  }

  if carryOver {
    if dn.Digit1 == 6 {
      // Shouldn't happen, due to higher level checks.
      panic("Overflow Error")
    } else {
      dn.Digit1++
    }
  }
}
