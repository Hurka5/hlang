package main

import (
  "os"
  "fmt"
  "sync"
  "errors"
  "io/ioutil"
  "lurka/internal/logger" 
  "lurka/internal/lexer" 
  "lurka/internal/parser" 
  "github.com/charmbracelet/log"
)



func init() {

  // TODO: better arg handling
  // Check for arg length 
  if len(os.Args) < 2 {
    os.Exit(1)
  }

  logger.StyleLogger()
}

func main()  {

  init() 
  
  files := []string{os.Args[1]} 

  var wg sync.WaitGroup
  wg.Add(len(files)) 

  errCollection := errors.New("")
  for _, file := range files {
      
    src, err := ioutil.ReadFile(file)
    if err != nil {
      log.Fatal(err.Error())
      wg.Done() 
      os.Exit(1)
    }

    // Lex
    l := lexer.New()
    tokens := l.Start(string(src))

    // Parse
    p := parser.New()
    p.Start(tokens)
    
    wg.Done()
    errCollection = errors.Join(l.Errors, p.Errors)
  }
  
  wg.Wait()
  log.Error(fmt.Sprintf("%s:%s", "asd", errCollection.Error()))

  // TODO: Wait for finish

  //if len(l.Errors) 

  // TODO: Generate asm

  // Assemble

  // Link
  
}
