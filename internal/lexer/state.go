package lexer

import (
  "fmt"
  "strings"
  "lurka/internal/token"
)

const (
  VALID_CHARS   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
  NUMS          = ""
  SPECIAL_CHARS = ""
)


type StateFunc func(*Lexer) StateFunc

func WhitespaceState(l *Lexer) StateFunc {
  r := l.Next()
	if r == EOF { return nil }

  // If New line found emit
	if r == '\n' {
    l.Emit(token.NL); 
    return WhitespaceState; 
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
  l.NewError(fmt.Sprintf("unexpected character: %q", r))
  l.Ignore()
  
  return WhitespaceState
}


func NumberState(l *Lexer) StateFunc {
  l.Take(NUMS)
  // TODO: Handle integers in octal (077) and hexadecimal(0x2F or 0XaD) forms
  // TODO: Handle float cases like: .1, 0.
  if l.Peek() == '.' {
    l.Take(NUMS)
    l.Emit(token.FLOAT_LIT)
  }else{
    l.Emit(token.INT_LIT)
  }
  return WhitespaceState
}


func IdentifierState(l *Lexer) StateFunc {
  l.Take(VALID_CHARS + NUMS)
  l.Emit(token.Lookup(l.Current()))
  return WhitespaceState
}


func SpecialState(l *Lexer) StateFunc {
  // TODO: Handle compley operators like: += or :=
  l.Take(SPECIAL_CHARS)
  l.Emit(token.Lookup(l.Current()))
  return WhitespaceState
}

func TextState(l *Lexer) StateFunc {
  // TODO: Handle string and char literals including escape characters
  return WhitespaceState
}
