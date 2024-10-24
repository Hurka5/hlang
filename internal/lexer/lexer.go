package lexer

import (
  "github.com/Hurka5/hlang/internal/token"
  "unicode/utf8"
  "strings"
  "fmt"
  "errors"
  _"github.com/charmbracelet/log"
)

type Lexer struct {
	source           string
	start, pos       int
	state            StateFunc
	tokens           chan token.Token
	Errors           error
	rewind           runeStack
}

const EOF rune = -1

// Creates and returns a lexer and a read only token channel
func New(src string) (*Lexer, <-chan token.Token) {
  
  l := &Lexer{
    source:     src,
    state:      WhitespaceState,
		rewind:     newRuneStack(),
    tokens:     make(chan token.Token, 100), // Buffer size - 100
	}
  go l.run()
  return l, l.tokens
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
  err := errors.New(fmt.Sprintf("%d:%d: %s",pos.Line, pos.Col, e))
  l.Errors = errors.Join(l.Errors, err)
}

// Private methods
func (l *Lexer) run() {
	state := l.state
	for state != nil {
		state = state(l)
	}
	close(l.tokens)
}
