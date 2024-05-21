package token


type TokenType int


type Token struct {
  Type TokenType 
  Literal string 
  Line int
  Col int
}


func Lookup(l string) TokenType {
  if kind, ok := keywords[l]; ok {
    return kind
  }
  return ID
}


func (t Token) String() string {
  for k, v :=  range  keywords {
    if(t.Type == v) {
      return k
    }
  }
  return ""
}


var keywords = map[string]TokenType{
  // Types
  "i32":    I32,
  "u32":    U32,
  "f32":    F32,
  "char":   CHAR,
  "string": STRING,
  "bool":   BOOL,

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
  "=": EQU,
  
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
  ERR TokenType = iota // unindentified characters
  NL                   // new line character
  ID                   // identifier 

  // Literals
  I32_LIT    // This is the type of a 32 bit signed integer value 
  U32_LIT    // This is the type of a 32 bit unsigned integer value
  F32_LIT    // This is the type of a 32 bit float value 
  CHAR_LIT   // This is the type of a 8 bit char value 
  STRING_LIT // This is the type of a string value
  TRUE      // true
  FALSE     // false

  // Keywords
  IF         // if      - if statement
  FOR        // for     - for loop
  FN         // fn      - function declaration
  WHILE      // while   - while loop
  PRINT      // print   - built in print function
  PRINTLN    // println - build in println function
  
  // Type Keywords
  I32       // 'i32' 
  U32       // 'u32'
  F32       // 'f32'
  CHAR      // 'char'
  STRING    // 'string'
  BOOL      // 'bool'

  // Operators
  ADD        // +
  SUB        // -
  MUL        // *
  DIV        // /
  MOD        // %
  EQU        // =
  
  // Seperators 
  DOT         // .
  COMMA       // ,
  COLON       // :
  SEMICOLON   // ;

  // Braces
  LPAR      // (
  RPAR      // )
  LBRACKET  // [
  RBRACKET  // ]
  LBRACE    // {
  RBRACE    // }

)


