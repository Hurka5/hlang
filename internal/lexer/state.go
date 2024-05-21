package lexer

import (
  "fmt"
  "errors"
  "strings"
  "hlang/internal/token"
  "hlang/internal/lexer/runestack"
)

const (
  VALID_CHARS   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
  NUMS          = ""
  SPECIAL_CHARS = ""
)


type StateFunc func(*Lexer) StateFunc

func DefaultState(l *Lexer) StateFunc {
  r := l.Next()
	if r == runestack.EOF { return nil }

  // If a new line found emit
	if r == '\n' {
    l.Emit(token.NL); 
    return DefaultState; 
  }

  l.Take(" \t\r")
	l.Ignore()

  // Check for numeric literals
  if strings.ContainsRune(NUMS, r)          { return NumberState; }

  // Check for special characters
  if strings.ContainsRune(SPECIAL_CHARS, r) { return SpecialState; }

  // Check for char or string literals
  if strings.ContainsRune(`'"`, r)          { return TextState; }

  // Check for keywords
  if strings.ContainsRune(VALID_CHARS, r)   { return IdentifierState; }

  // Unexpected Character
  //l.NewError(fmt.Sprintf("unexpected character: %q", r))
  errors.Join(l.Errors, errors.New(fmt.Sprintf("unexpected character: %q", r)))
  l.Ignore()
  
  return DefaultState
}


func NumberState(l *Lexer) StateFunc {
  l.Take(NUMS)
  // TODO: Handle integers in octal (077) and hexadecimal(0x2F or 0XaD) forms
  // TODO: Handle float cases like: .1, 0.
  if l.Peek() == '.' {
    l.Take(NUMS)
    l.Emit(token.F32_LIT)
  }else{
    l.Emit(token.I32_LIT)
  }
  return DefaultState
}


func IdentifierState(l *Lexer) StateFunc {
  l.Take(VALID_CHARS + NUMS)
  l.Emit(token.Lookup(l.Current()))
  return DefaultState
}


func SpecialState(l *Lexer) StateFunc {
  // TODO: Handle complex operators like: += or :=
  l.Take(SPECIAL_CHARS)
  l.Emit(token.Lookup(l.Current()))
  return DefaultState
}

func TextState(l *Lexer) StateFunc {
  // TODO: Handle string and char literals including escape characters
  return DefaultState
}
