package lexer

import (
  "testing"
  "hlang/internal/token"
)

func TestLexerEmpty(t *testing.T) {
  l := New()
  ch := l.StartWithSource("");
  for tok := range ch {
    t.Fatalf("Found %s, expected: none", tok.String())
  }
}

func TestLexerUnicode(t *testing.T) {
  l := New()
  ch := l.StartWithSource("á");
  for tok := range ch {
    if tok.Type != token.ERR {
      t.Fatalf("Found %s, expected: ERR", tok.String())
    }
  }
}

var testCases =  map[string][]token.TokenType {
  "u32 asd": []token.TokenType{ token.U32 },
  "i32 asd = -12": []token.TokenType{ token.I32, token.ID, token.EQU, token.I32_LIT },
  "f32  asd = 0.12": []token.TokenType{ token.I32, token.ID, token.EQU, token.I32_LIT },
  "f32 asd = -3.14": []token.TokenType{ token.I32, token.ID, token.EQU, token.I32_LIT },
  "bool asd, dsa = false, true ": []token.TokenType{ token.I32, token.ID, token.EQU, token.I32_LIT },
}

func TestLexerReading(t *testing.T) {
  for k, v := range testCases {
  l := New()
    ch := l.StartWithSource(k);
    for tok := range ch {
      if(len(v) == 0){ 
        t.Fatalf("Channel contains more token: %s", tok.String())
      }
      if(tok.Type == v[0]) {
        t.Fatalf("Channel contains more")
        v = v[1:]
      }
    }
  }

}
func TestLexerErrorCatch(t *testing.T) {}
