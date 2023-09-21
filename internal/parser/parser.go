package parser

import (
  "errors"
  "lurka/internal/logger"
  "lurka/internal/token"
  "lurka/internal/ast"
)

type Parser struct {
  Errors   error
  state    StateFunc
  tokens   *chan token.Token
  rootNode ast.Node
}

// Walker - a object that goes throught the tokens

func New() *Parser {
  return &Parser{
    state: nil,
  }
}

func (p *Parser) Start(toks *chan token.Token) *ast.Node {
  p.tokens = toks
  go p.run()
  return &p.rootNode
}

// NextToken returns the next token from the channel and a value to denote whether
// or not the token is finished.
func (p *Parser) NextToken() (*token.Token, bool) {
	if tok, ok := <- *p.tokens; ok {
		return &tok, false
	} else {
		return nil, true
	}
}

func (p *Parser) getPos(e string) {
  
}


func (p *Parser) NewError(e string) {
  //pos := p.getPos()
  p.Errors = errors.Join(p.Errors, logger.SyntaxError{
    Line: 0, // CHANGE THIS
    Col:  0,  // CHANGE THIS
    Msg:  e,
  })
}
  
// Private methods
func (p *Parser) run() {
	state := p.state
	for state != nil {
		state = state(p)
	}
}
