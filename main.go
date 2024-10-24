package main

import (
  "os"
  "github.com/charmbracelet/log"
  "github.com/charmbracelet/lipgloss"
  "github.com/Hurka5/hlang/internal/lexer"
  _"errors"
  "fmt"
  "io/ioutil"
)

var logger *log.Logger

func init() {

  // Set styles
  styles := log.DefaultStyles()

  // ERROR Style
  styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
    SetString("ERROR").
	  Padding(0, 1, 0, 1).
	  Background(lipgloss.Color("204")).
	  Foreground(lipgloss.Color("0"))

  // WARNING Style
  styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
    SetString("WARNING").
	  Padding(0, 1, 0, 1).
	  Background(lipgloss.Color("11")).
	  Foreground(lipgloss.Color("0"))

  // INFO Style
  styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
    SetString("INFO").
	  Padding(0, 1, 0, 1).
	  Background(lipgloss.Color("43")).
	  Foreground(lipgloss.Color("0"))

  // Create logger
  logger = log.NewWithOptions(os.Stderr, log.Options{
    ReportCaller: false,
    ReportTimestamp: false,
    Prefix: "hlang",
  })

  // Assign styles
  logger.SetStyles(styles)
}

func main() {

  // Check if we got input
  if len(os.Args) < 2 {
    err := fmt.Errorf("no input file")
    logger.Warn(err)
    os.Exit(1)
  }
  
  // Save file name
  in_file := os.Args[1] 

  // Read code out of file
  src, err := ioutil.ReadFile(in_file)
  if err != nil {
    fmt.Fprintf(os.Stderr,err.Error())
    os.Exit(1)
  }

  // Create lexer
  l, tokens := lexer.New(string(src));
  

  for t := range tokens {
    fmt.Println(t.Kind)
  }

  logger.Error(l.Errors)
}
