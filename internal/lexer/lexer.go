package lexer

import (
	"errors"
	"strings"
	"unicode/utf8"
  "lurka/internal/token"
  "lurka/internal/logger"
)

const EOF rune = -1

type Lexer struct {
	source           string
	start, pos       int
	state            StateFunc
	tokens           chan token.Token
	Errors           error
	rewind           runeStack
}

// New creates a returns a lexer ready to parse the given source code.
func New() *Lexer {
	return &Lexer{
    state:      WhitespaceState,
		start:      0,
		pos:        0,
		rewind:     newRuneStack(),
	}
}

// Start begins executing the Lexer in an asynchronous manner (using a goroutine)
// and returns the channel where the tokens stored
func (l *Lexer) Start(src string) <- chan token.Token {

  l.source = src

	l.tokens = make(chan token.Token, 100) // buffer size - 100 

	go l.run()

  return l.tokens
}

func (l *Lexer) Tokens() <- chan token.Token {
  tokens := make(chan token.Token, 0)
  go func() {
      for state != nil {
        tok := s.Scan()
      }
  }()
  return tokens
}


// Current returns the value being being analyzed at this moment.
func (l *Lexer) Current() string {
	return l.source[l.start:l.pos]
}

// Emit will receive a token type and push a new token with the current analyzed
// value into the tokens channel.
func (l *Lexer) Emit(k token.TokenKind) {
	tok := token.Token{
		Kind:    k,
		Literal: l.Current(),
    Pos:     l.getPos(),
	}
  
	l.tokens <- tok
	l.start = l.pos
	l.rewind.clear()
}

// Ignore clears the rewind stack and then sets the current beginning pos
// to the current pos in the source which effectively ignores the section
// of the source being analyzed.
func (l *Lexer) Ignore() {
	l.rewind.clear()
	l.start = l.pos
}

// Peek performs a Next operation immediately followed by a Rewind returning the
// peeked rune.
func (l *Lexer) Peek() rune {
	r := l.Next()
	l.Rewind()

	return r
}

// Rewind will take the last rune read (if any) and rewind back. Rewinds can
// occur more than once per call to Next but you can never rewind past the
// last point a token was emitted.
func (l *Lexer) Rewind() {
	r := l.rewind.pop()
	if r > EOF {
		size := utf8.RuneLen(r)
		l.pos -= size
		if l.pos < l.start {
			l.pos = l.start
		}
	}
}

// Next pulls the next rune from the Lexer and returns it, moving the pos
// forward in the source.
func (l *Lexer) Next() rune {
	var (
		r rune
		s int
	)
	str := l.source[l.pos:]
	if len(str) == 0 {
		r, s = EOF, 0
	} else {
		r, s = utf8.DecodeRuneInString(str)
	}
	l.pos += s
	l.rewind.push(r)

	return r
}

// Take receives a string containing all acceptable strings and will contine
// over each consecutive character in the source until a token not in the given
// string is encountered. This should be used to quickly pull token parts.
func (l *Lexer) Take(chars string) {
	r := l.Next()
	for strings.ContainsRune(chars, r) {
		r = l.Next()
	}
	l.Rewind() // last next wasn't a match
}

func (l *Lexer) getPos() token.Position {
	untilNow := l.source[:l.pos]
  
	lineNum := strings.Count(untilNow, "\n") + 1
	lastNewLineIndex := strings.LastIndex(untilNow, "\n")
	posInLine := l.pos - lastNewLineIndex

  return token.Position{lineNum, posInLine}
}

// Partial yyLexer implementation
func (l *Lexer) NewError(e string) {
  pos := l.getPos()
  l.Errors = errors.Join(l.Errors, logger.SyntaxError{
    Line: pos.Line,
    Col:  pos.Col,
    Msg:  e,
  })
}

// Private methods
func (l *Lexer) run() {
	state := l.state
	for state != nil {
		state = state(l)
	}
	close(l.tokens)
}
