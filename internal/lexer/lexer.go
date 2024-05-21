package lexer


import (
  "os"
  "errors"
  "strings"
  "unicode/utf8"
  "hlang/internal/token"
  "hlang/internal/lexer/runestack"
)


func New() *Lexer {
	return &Lexer{
    state:      DefaultState,
		start:      0,
		pos:        0,
    tokens:     make(chan token.Token, 100), // channel buffer size is 101 
	}
}

type Lexer struct {
	source           string
	start, pos       int
	state            StateFunc
	tokens           chan token.Token
  runeStack        runestack.RuneStack
	Errors           error
}

func (l *Lexer) StartWithSource(source string) <-chan token.Token {
  l.source = source

  // Start the lexical analisis
  go l.run()

  return l.tokens
}


func (l *Lexer) Start(filename string) <-chan token.Token {

  // Read all file and store it in source
  src, err := os.ReadFile(filename)
  if err != nil {
    errors.Join(l.Errors, err)
  }
  l.source = string(src)

  // Start the lexical analisis
  go l.run()

  return l.tokens
}


// Next() pulls the next rune from the Lexer and returns it, moving the pos forward in the source.
func (l *Lexer) Next() rune {
	var (
		r rune
		s int
	)
	str := l.source[l.pos:]
	if len(str) == 0 {
		r, s = runestack.EOF, 0
	} else {
		r, s = utf8.DecodeRuneInString(str)
	}
	l.pos += s
	l.runeStack.Push(r)

	return r
}


/*
  Take receives a string containing all acceptable strings and will contine
  over each consecutive character in the source until a rune is not in the given
  string is encountered. This should be used to quickly pull token parts. 
*/
func (l *Lexer) Take(chars string) {
	r := l.Next()
	for strings.ContainsRune(chars, r) {
		r = l.Next()
	}
	l.Rewind() // Last next wasn't a match
}


/* 
  Rewind() will take the last rune read (if any) and rewind back. Rewinds can
  occur more than once per call to Next() but you can never rewind past the
  last point a token was emitted.
*/
func (l *Lexer) Rewind() {
	r := l.runeStack.Pop()
	if r > runestack.EOF {
		size := utf8.RuneLen(r)
		l.pos -= size
		if l.pos < l.start {
			l.pos = l.start
		}
	}
}


/*
  Emit will receive a token type and push a new token with the current analyzed
  value into the tokens channel. 
*/
func (l *Lexer) Emit(k token.TokenType) {

	untilNow := l.source[:l.pos]
  
  //TODO: Change this 
	line := strings.Count(untilNow, "\n") + 1
	lastNewLineIndex := strings.LastIndex(untilNow, "\n")
	col := l.pos - lastNewLineIndex

  tok := token.Token{
		Type:    k,
		Literal: l.Current(),
    Line:    line,
    Col:     col,
	}

	l.tokens <- tok
	l.start = l.pos
	l.runeStack.Clear()
}


// Peek performs a Next operation immediately followed by a Rewind returning the peeked rune.
func (l *Lexer) Peek() rune {
	r := l.Next()
	l.Rewind()

	return r
}


// Current returns the value being being analyzed at this moment.
func (l *Lexer) Current() string {
	return l.source[l.start:l.pos]
}


/* 
  Ignore clears the rune stack and then sets the current beginning pos
  to the current pos in the source which effectively ignores the section
  of the source being analyzed. 
*/
func (l *Lexer) Ignore() {
	l.runeStack.Clear()
	l.start = l.pos
}








// run() is the real loop of the lexer it is called from the start it cannot occure multiple times 
func (l *Lexer) run() {
	state := l.state
	for state != nil {
		state = state(l)
	}
	close(l.tokens)
}

