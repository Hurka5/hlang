package main

import(
  "os"
  "sync"
  "hlang/internal/lexer"
)

func main() {
  
  files := os.Args[1:]
  
  var wg sync.WaitGroup 

  for _, filename := range files {
    go func(){
      wg.Add(1)
      lex := lexer.New();
      lex.Start(filename)
      //pars := Parser.New(tokens)
      //_ := pars.Start() // AST
      wg.Done()
    }()
  }
  wg.Wait()
  

}
