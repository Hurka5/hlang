package lexer

import (
  "testing"
  "github.com/Hurka5/hlang/internal/token"
)

func TryLex(input string) token.Token {
  _, tokens := New(input)
  return <- tokens
}

var NumberTests = map[string]token.TokenKind{
  ".12":  token.FLOAT,
  "0.12": token.FLOAT,
  "123":  token.INT,
  "0b1":  token.INT,
  "0o12": token.INT,
  "001":  token.INT,
  "0x12": token.INT,
  "asd":  token.ID,
  "for":  token.FOR,
  "(":    token.LPAR,
}
func TestNumberLexing(t *testing.T) {
  for in, kind := range NumberTests {
    tok := TryLex(in)
    if tok.Kind != kind {
      t.Fatalf("Lexing %q should be %q, instead its %q (%q)", in, kind, tok.Kind, tok.Literal) 
    }
  }
}
