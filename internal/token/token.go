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

var keywords = map[string]TokenKind{
  // Types
  "i32":    INT,
  "f32":    FLOAT,
  "char":   CHAR,
  "string": STRING,

  // Keywords
  "if":    IF,
  "for":   FOR,
  "fn":    FN,
  "while": WHILE,
  "print": PRINT,
  "println": PRINTLN,

  // Operators
  "+": ADD,
  "-": SUB,
  "*": MUL,
  "/": DIV,
  "%": MOD,
  
  // Seperators 
  ".": DOT,
  ",": COMMA,
  ":": COLON,
  ";": SEMICOLON,

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
  ERR TokenKind = iota
  NL 
  ID

  // Types
  INT
  FLOAT
  CHAR
  STRING

  // Literals
  INT_LIT
  FLOAT_LIT
  CHAR_LIT
  STRING_LIT
  BOOL_LIT

  // Keywords
  IF
  FOR 
  FN
  WHILE
  PRINT
  PRINTLN

  // Operators
  ADD
  SUB
  MUL
  DIV
  MOD
  
  // Seperators 
  DOT
  COMMA
  COLON
  SEMICOLON

  // Braces
  LPAR      // (
  RPAR      // )
  LBRACKET  // [
  RBRACKET  // ]
  LBRACE    // {
  RBRACE    // }

)

