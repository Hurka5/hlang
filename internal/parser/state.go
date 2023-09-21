package parser

type StateFunc func(*Parser) StateFunc

func MainState(p *Parser) StateFunc {

  return nil
}

