package lexer

import (
  "github.com/Hurka5/hlang/internal/token"
  "fmt"
  "strings"
)

const (
  VALID_CHARS   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
  NUMS          = "0123456789"
  SPECIAL_CHARS = "(){}[]"
  WHITESPACE    = " \t\r"
)

type StateFunc func(*Lexer) StateFunc

func WhitespaceState(l *Lexer) StateFunc {
  l.Take(WHITESPACE)
	l.Ignore()

  r := l.Next()
  nr := l.Peek()

	if r == EOF {
    return nil 
  }

  // If New line found emit
	if r == '\n' {
    l.Emit(token.NL); 
    return WhitespaceState; 
  }

  // Check for numeric literals
  if strings.ContainsRune(NUMS, r) || (r == '.' && strings.ContainsRune(NUMS, nr)) { return NumberState; }

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

  // Get first character
  l.Rewind()
  r := l.Next()
  wasZero := r == '0'
  
  // handling float numbers with '.' start
  if r == '.' {
    l.Take(NUMS)
    l.Emit(token.FLOAT)
    return WhitespaceState
  }

  l.Take(NUMS)

  // handling float numbers
  if l.Peek() == '.' {
    l.Take(NUMS)
    l.Emit(token.FLOAT)
    return WhitespaceState
  }
  
  r = l.Next()
  // handling hexa and octal and binary numbers
  if wasZero && strings.ContainsRune("xXoObB",r) && strings.ContainsRune(NUMS,l.Peek()) {
    // if binary
    if strings.ContainsRune("bB",r) {
      l.Take("01")
    }else{
      l.Take(NUMS)
    }
  }else{
    l.Rewind()
  }


  l.Emit(token.INT)

  return WhitespaceState
}


func IdentifierState(l *Lexer) StateFunc {
  l.Take(VALID_CHARS + NUMS)
  l.Emit(token.Lookup(l.Current()))
  return WhitespaceState
}


func SpecialState(l *Lexer) StateFunc {
  l.Emit(token.Lookup(l.Current()))
  // TODO: Handle complex operators like: += or :=
  return WhitespaceState
}

func TextState(l *Lexer) StateFunc {
  // TODO: Handle string and char literals including escape characters
  return WhitespaceState
}


