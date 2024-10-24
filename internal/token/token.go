package token

type Position struct {
	Line int
	Col  int
}

type TokenKind int

type Token struct {
  Kind    TokenKind
  Literal string
  Pos     Position
}

func Lookup(l string) TokenKind {
  if kind, ok := keywords[l]; ok {
    return kind
  }
  return ID
}

func (k TokenKind) String() string {
  for str, kind := range keywords {
    if k == kind{
      return str
    }
  }
  return "???"
}

var keywords = map[string]TokenKind{
  "ILLEGAL":  ILLEGAL,
  "NEW LINE": NL,
  "ID":       ID,
  "TYPE":     TYPE,
  // Types
  "int":    INT,
  "float":  FLOAT,
  "char":   CH,
  "string": STR,

  // Keywords
  "if":    IF,
  "for":   FOR,
  "fn":    FN,
  "while": WHILE,
  "print": PRINT,

  // Operators
  "+": PLUS,
  "-": MINUS,
  "*": MUL,
  "/": DIV,
  "%": MOD,
  
  // Braces
  "(": LPAR,    
  ")": RPAR,    
  "[": LBRACKET,
  "]": RBRACKET,
  "{": LBRACE,  
  "}": RBRACE,  
} 

const (
  // General
  ILLEGAL TokenKind = iota
  NL 
  ID
  TYPE

  // Types
  INT
  FLOAT
  CH
  STR

  // Keywords
  IF
  FOR 
  FN
  WHILE
  PRINT

  // Operators
  PLUS
  MINUS
  MUL
  DIV
  MOD
  
  // Braces
  LPAR      // (
  RPAR      // )
  LBRACKET  // [
  RBRACKET  // ]
  LBRACE    // {
  RBRACE    // }

)

